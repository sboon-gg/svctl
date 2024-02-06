package server

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/pkg/errors"
	"github.com/sboon-gg/svctl/pkg/templates"
)

const (
	SvctlDir     = ".svctl"
	TemplatesDir = "templates"

	ConfigFile  = "config.yaml"
	ValuesFile  = "values.yaml"
	MaplistFile = "maplist.yaml"

	CacheFile = ".cache.yaml"

	maxRestartRetries = 10
)

type Server struct {
	log *slog.Logger

	Path    string
	process *os.Process

	restartCancel context.CancelFunc
}

func Open(serverPath string) (*Server, error) {
	_, err := os.Stat(filepath.Join(serverPath, SvctlDir))
	if err != nil {
		return nil, errors.Wrap(err, "server not initialized")
	}

	s := &Server{
		log:  slog.Default().With(slog.String("serverPath", serverPath)),
		Path: serverPath,
	}

	cache, err := s.Cache()
	if err != nil {
		return nil, err
	}

	if cache.PID != -1 {
		s.process, err = os.FindProcess(cache.PID)
		if err != nil {
			err = s.unsetProcess()
			if err != nil {
				return nil, err
			}
		}

		if err = healthCheck(s.process); err != nil {
			err = s.unsetProcess()
			if err != nil {
				return nil, err
			}
		}

		if s.process != nil {
			// TODO: verify that no one else is managing the process
			s.setupRestart()
		}
	}

	return s, nil
}

type Opts struct {
	TemplatesRepo string
	Token         string
}

func Initialize(serverPath string, opts *Opts) (*Server, error) {
	if opts == nil {
		opts = &Opts{}
	}

	svctlPath := filepath.Join(serverPath, SvctlDir)
	err := os.Mkdir(svctlPath, 0755)
	if err != nil {
		return nil, err
	}

	err = writeDefaultConfig(svctlPath)
	if err != nil {
		return nil, err
	}

	if opts.TemplatesRepo != "" {
		templatesPath := filepath.Join(svctlPath, TemplatesDir)
		err = cloneTemplates(templatesPath, opts.TemplatesRepo, opts.Token)
		if err != nil {
			return nil, err
		}

		t, err := templates.NewFromPath(templatesPath)
		if err != nil {
			return nil, err
		}

		err = writeValues(svctlPath, t.DefaultsContent())
		if err != nil {
			return nil, err
		}
	}

	return Open(serverPath)
}

func (s *Server) Start() error {
	if s.process != nil {
		return nil
	}

	process, err := startProcess(s.Path)
	if err != nil {
		return err
	}

	err = s.setProcess(process)
	if err != nil {
		return err
	}

	s.setupRestart()

	return nil
}

func (s *Server) Stop() error {
	if s.process == nil {
		return nil
	}

	s.restartCancel()

	if err := stopProcess(s.process); err != nil {
		return err
	}

	return s.unsetProcess()
}

func (s *Server) setProcess(proc *os.Process) error {
	s.process = proc

	return s.storePID(proc.Pid)
}

func (s *Server) unsetProcess() error {
	s.process = nil

	return s.storePID(-1)
}

func (s *Server) setupRestart() {
	restartCtx, restartCancel := context.WithCancel(context.Background())
	s.restartCancel = restartCancel

	processExited := false

	wait := func() {
		if !processExited {
			_, err := s.process.Wait()
			if err != nil {
				s.log.Error(fmt.Sprintf("Failed to wait for process: %s", err))
			}

			processExited = true
		}
	}

	go func() {
		go wait()
		retries := 0
		for {
			select {
			case <-restartCtx.Done():
				return
			default:
				if err := healthCheck(s.process); err != nil || processExited {
					retries++
					if retries >= maxRestartRetries {
						s.log.Error(fmt.Sprintf("Failed to restart process after %d retries, giving up", maxRestartRetries))
						s.restartCancel()
						return
					}

					err := s.restart()
					if err != nil {
						s.log.Error(fmt.Sprintf("Failed to restart process: %s", err))
						return
					}

					retries = 0
					processExited = false
					go wait()
				}
			}

			// time.Sleep(500 * time.Millisecond)
			time.Sleep(2 * time.Second)
		}
	}()
}

func (s *Server) restart() error {
	err := stopProcess(s.process)
	if err != nil {
		return errors.Wrap(err, "Failed to gracefully shutdown process before restart")
	}

	err = s.unsetProcess()
	if err != nil {
		return err
	}

	proc, err := startProcess(s.Path)
	if err != nil {
		return errors.Wrap(err, "Failed to restart process")
	}

	return s.setProcess(proc)
}
