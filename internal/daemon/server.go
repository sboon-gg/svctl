package daemon

import (
	"log/slog"
	"os"

	"github.com/sboon-gg/svctl/internal/daemon/fsm"
	"github.com/sboon-gg/svctl/internal/server"
)

type RunningServer struct {
	server.Server

	fsm *fsm.FSM
	log *slog.Logger
}

func OpenServer(svPath string) (*RunningServer, error) {
	s, err := server.Open(svPath)
	if err != nil {
		return nil, err
	}

	rs := &RunningServer{
		Server: *s,
		log:    initLog(svPath),
	}

	rs.fsm = fsm.New(svPath, rs.setProcess)

	cache, err := s.Cache()
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

func (rs *RunningServer) setProcess(_ fsm.StateT) {
	rs.log.Info("Setting process")

	pid := -1
	if rs.fsm != nil {
		pid = rs.fsm.Pid()
	}

	rs.log = initLog(rs.Path).With(slog.Int("pid", pid))

	err := rs.StorePID(pid)
	if err != nil {
		rs.log.Error("Failed to store PID", slog.String("err", err.Error()))
	}
}

func initLog(svPath string) *slog.Logger {
	return slog.Default().With(slog.String("sv", svPath))
}
