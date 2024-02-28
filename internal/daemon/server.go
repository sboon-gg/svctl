package daemon

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/sboon-gg/svctl/internal/daemon/fsm"
	"github.com/sboon-gg/svctl/internal/server"
	"github.com/sboon-gg/svctl/internal/settings"
)

type RunningServer struct {
	server.Server

	fsm *fsm.FSM
	log *slog.Logger
}

func OpenServer(svPath string) (*RunningServer, error) {
	settingsPath := filepath.Join(svPath, settings.SvctlDir)
	s, err := server.Open(svPath, settingsPath)
	if err != nil {
		return nil, err
	}

	rs := &RunningServer{
		Server: *s,
		log:    initLog(svPath),
	}

	rs.fsm = fsm.New(svPath, rs.setProcess, rs.fullRender)

	cache, err := s.Settings.Cache()
	if err != nil {
		return nil, err
	}

	if cache.PID != -1 {
		proc, err := os.FindProcess(cache.PID)
		if err == nil {
			err = rs.fsm.Adopt(proc)
			if err != nil {
				return nil, err
			}
		}
	}

	return rs, nil
}

func (rs *RunningServer) Start() error {
	rs.log.Info("Rendering templates")

	err := rs.fullRender()
	if err != nil {
		return err
	}

	rs.log.Info("Starting server")

	return rs.fsm.Start()
}

func (rs *RunningServer) Stop() error {
	rs.log.Info("Stopping server")

	return rs.fsm.Stop()
}

func (rs *RunningServer) Restart() error {
	rs.log.Info("Restarting server")

	return rs.fsm.Restart()
}

func (rs *RunningServer) setProcess(newState fsm.StateT) {
	if newState == fsm.StateTRestarting {
		rs.log.Info("Process restarting")
	}

	if newState == fsm.StateTStopped {
		rs.log.Info("Process stopped")
	}

	rs.log.Debug("Setting PID in cache")

	pid := -1
	if rs.fsm != nil {
		pid = rs.fsm.Pid()
	}

	rs.log = initLog(rs.ServerPath).With(slog.Int("pid", pid))

	err := rs.Settings.StorePID(pid)
	if err != nil {
		rs.log.Error("Failed to store PID", slog.String("err", err.Error()))
	}
}

func (rs *RunningServer) fullRender() error {
	outputs, err := rs.Render()
	if err != nil {
		return err
	}

	return rs.WriteTemplatesOutput(outputs)
}

func initLog(svPath string) *slog.Logger {
	return slog.Default().With(slog.String("sv", svPath))
}
