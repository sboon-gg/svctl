package daemon

import (
	"os"
	"path/filepath"

	"github.com/sboon-gg/svctl/internal/daemon/fsm"
	"github.com/sboon-gg/svctl/internal/server"
	"github.com/sboon-gg/svctl/internal/settings"
)

func OpenServer(svPath string) (*fsm.FSM, error) {
	settingsPath := filepath.Join(svPath, settings.SvctlDir)
	s, err := server.Open(svPath, settingsPath)
	if err != nil {
		return nil, err
	}

	svFSM := fsm.New(s)

	cache, err := s.Settings.Cache()
	if err != nil {
		return nil, err
	}

	if cache.PID != -1 {
		proc, err := os.FindProcess(cache.PID)
		if err == nil {
			err = svFSM.Adopt(proc)
			if err != nil {
				// Process can be already dead
				cache.PID = -1
				err = s.Settings.WriteCache(cache)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return svFSM, nil
}
