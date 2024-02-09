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

func initLog(svPath string) *slog.Logger {
	return slog.Default().With(slog.String("sv", svPath))
}

func Open(serverPath string) (*Server, error) {
	_, err := os.Stat(filepath.Join(serverPath, SvctlDir))
	if err != nil {
		return nil, errors.Wrap(err, "server not initialized")
	}

	s := &Server{
		log:  initLog(serverPath),
		Path: serverPath,
	}

	cache, err := s.Cache()
	if err != nil {
		return nil, err
	}

	if cache.PID != -1 {
		s.process, err = os.FindProcess(cache.PID)
		if err != nil || !isHealthy(s.process) {
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
	s.log.Info("Starting server")

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
	s.log.Info("Stopping server")

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
	s.log.Info("Setting process")

	s.process = proc

	s.log = initLog(s.Path).With(slog.Int("pid", proc.Pid))

	return s.storePID(proc.Pid)
}

func (s *Server) unsetProcess() error {
	s.log.Info("Unsetting process")

	s.process = nil

	s.log = initLog(s.Path).With(slog.Int("pid", -1))

	return s.storePID(-1)
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
		return errors.Wrap(err, "Failed to start new process")
	}

	return s.setProcess(proc)
}

func (s *Server) setupRestart() {
	restartCtx, restartCancel := context.WithCancel(context.Background())
	s.restartCancel = restartCancel

	r := &restarter{
		s:             s,
		maxRetries:    maxRestartRetries,
		restartCtx:    restartCtx,
		restartCancel: restartCancel,
	}

	go r.Watch()
}

type restarter struct {
	s *Server

	retries    int
	maxRetries int

	restartCtx    context.Context
	restartCancel context.CancelFunc
}

func (r *restarter) Watch() {
	go r.attemptToWait()

	for {
		select {
		case <-r.restartCtx.Done():
			return
		default:
			if isHealthy(r.s.process) {
				r.retries = 0
				time.Sleep(500 * time.Millisecond)
				continue
			}

			err := r.attempRestart()
			if err != nil {
				r.restartCancel()
				r.s.log.Error(err.Error())
				return
			}

			go r.attemptToWait()
		}
	}
}

func (r *restarter) attempRestart() error {
	if r.retries >= r.maxRetries {
		return fmt.Errorf("Failed to restart process after %d retries, giving up", r.maxRetries)
	}

	r.retries++

	err := r.s.restart()
	if err != nil {
		r.s.log.Error(fmt.Sprintf("Failed to restart process: %s", err))
		return r.attempRestart()
	}

	r.retries = 0

	return nil
}

// attemptToWait waits for the process to exit and logs the state
// it might be unable to wait if it didn't start the process (daemon recover)
// or the process might be already dead
// release the process so the health check in Watch() can detect it
func (r *restarter) attemptToWait() {
	state, err := r.s.process.Wait()
	if err == nil {
		r.s.log.Info(fmt.Sprintf("Process exited with state: %s", state.String()))
		_ = r.s.process.Release()
	}
}
