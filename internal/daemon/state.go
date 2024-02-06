package daemon

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type State struct {
	Servers []string
}

func NewState() *State {
	return &State{
		Servers: make([]string, 0),
	}
}

func (d *Daemon) State() (*State, error) {
	var cache State

	if _, err := os.Stat(d.cachePath(stateFile)); os.IsNotExist(err) {
		return NewState(), nil
	}

	content, err := os.ReadFile(d.cachePath(stateFile))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &cache)
	if err != nil {
		return nil, err
	}

	return &cache, nil
}

func (d *Daemon) SaveState(state *State) error {
	content, err := yaml.Marshal(state)
	if err != nil {
		return err
	}

	return os.WriteFile(d.cachePath(stateFile), content, 0644)
}

func (d *Daemon) cachePath(path string) string {
	return filepath.Join(d.cacheDir, path)
}
