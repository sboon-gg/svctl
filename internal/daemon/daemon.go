package daemon

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sboon-gg/svctl/internal/daemon/fsm"
)

const (
	svctlDir  = "svctl"
	stateFile = "state.yaml"
)

type Daemon struct {
	cacheDir string
	Servers  map[string]*fsm.FSM
}

func New() (*Daemon, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	svctlCacheDir := filepath.Join(cacheDir, svctlDir)

	err = os.MkdirAll(svctlCacheDir, 0755)
	if err != nil {
		return nil, err
	}

	return &Daemon{
		Servers:  make(map[string]*fsm.FSM),
		cacheDir: svctlCacheDir,
	}, nil
}

func Recover() (*Daemon, error) {
	d, err := New()
	if err != nil {
		return nil, err
	}

	state, err := d.State()
	if err != nil {
		return nil, err
	}

	for _, svPath := range state.Servers {
		s, err := OpenServer(svPath)
		if err != nil {
			return nil, err
		}

		d.Servers[svPath] = s
	}

	return d, nil
}

func (s *Daemon) Register(path string) error {
	if _, ok := s.Servers[path]; ok {
		return fmt.Errorf("server %q already exists", path)
	}

	sv, err := OpenServer(path, s.updaterCache)
	if err != nil {
		return err
	}

	s.Servers[path] = sv

	state, err := s.State()
	if err != nil {
		return err
	}

	state.Servers = append(state.Servers, path)

	return s.SaveState(state)
}

func (s *Daemon) Start(path string) error {
	srv, err := s.findServer(path)
	if err != nil {
		return err
	}

	return srv.Start()
}

func (s *Daemon) Stop(path string) error {
	srv, err := s.findServer(path)
	if err != nil {
		return err
	}
	return srv.Stop()
}

func (d *Daemon) findServer(path string) (*fsm.FSM, error) {
	s, ok := d.Servers[path]
	if !ok {
		return nil, fmt.Errorf("server %q not found", path)
	}

	return s, nil
}
