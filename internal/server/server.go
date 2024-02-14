package server

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sboon-gg/svctl/internal/server/prbf2"
	"github.com/sboon-gg/svctl/pkg/templates"
)

const (
	SvctlDir     = ".svctl"
	TemplatesDir = "templates"

	ConfigFile  = "config.yaml"
	ValuesFile  = "values.yaml"
	MaplistFile = "maplist.yaml"

	CacheFile = ".cache.yaml"
)

type Server struct {
	log *slog.Logger

	Path string
	proc *prbf2.PRBF2
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

	s.proc = prbf2.New(serverPath, s.setProcess)

	cache, err := s.Cache()
	if err != nil {
		return nil, err
	}

	if cache.PID != -1 {
		proc, err := os.FindProcess(cache.PID)
		if err == nil {
			err = s.proc.Adopt(proc)
			if err != nil {
				return nil, err
			}
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

	return s.proc.Start()
}

func (s *Server) Stop() error {
	s.log.Info("Stopping server")

	return s.proc.Stop()
}

func (s *Server) Restart() error {
	s.log.Info("Restarting server")

	return s.proc.Restart()
}

func (s *Server) setProcess(_ prbf2.State) {
	s.log.Info("Setting process")

	pid := -1
	if s.proc != nil {
		pid = s.proc.Pid()
	}

	s.log = initLog(s.Path).With(slog.Int("pid", pid))

	err := s.storePID(pid)
	if err != nil {
		s.log.Error("Failed to store PID", slog.String("err", err.Error()))
	}
}
