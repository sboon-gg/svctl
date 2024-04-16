package fsm

import (
	"context"
	"errors"
	"log/slog"
	"time"
)

type stateEmpty struct{}

func (s *stateEmpty) Enter(fsm *FSM) {}
func (s *stateEmpty) Exit()          {}

type StateStopped struct {
	stateEmpty
}

func (s *StateStopped) Enter(fsm *FSM) {
	log := fsm.server.Settings.Log.With(slog.String("state", "stopped"))

	log.Debug("Stopping server")

	err := fsm.server.Settings.StorePID(-1)
	if err != nil {
		log.Error("Failed to store PID", "error", err.Error())
	}

	err = fsm.proc.Stop()
	if err != nil {
		fsm.handleError(err)
	}

	log.Info("Server stopped")

	fsm.cancel()
}

type StateRunning struct {
	cancel context.CancelFunc
}

func (s *StateRunning) Enter(fsm *FSM) {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	log := fsm.server.Settings.Log.With(slog.String("state", "running"))

	if !fsm.proc.IsRunning() {
		log.Info("Rendering templates")
		err := fsm.server.Render()
		if err != nil {
			fsm.handleError(err)
			return
		}

		log.Info("Starting server")
		err = fsm.proc.Start()
		if err != nil {
			fsm.handleError(err)
			return
		}
	}

	pid := fsm.proc.Pid()

	log = log.With(slog.Int("pid", pid))

	err := fsm.server.Settings.StorePID(pid)
	if err != nil {
		log.Error("Failed to store PID", "error", err.Error())
	}

	go func() {
		fsm.proc.Wait()
		log.Debug("Process exited")

		cancel()
		fsm.ChangeState(StateTRestarting)
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := fsm.server.Render()
				if err != nil {
					log.Error(errors.Join(errors.New("Failed to render templates"), err).Error())
				}
				time.Sleep(1 * time.Minute)
			}
		}
	}()
}

func (s *StateRunning) Exit() {
	s.cancel()
}

type StateRestarting struct {
	stateEmpty
	restartCtx restarter
}

func (s *StateRestarting) Enter(fsm *FSM) {
	log := fsm.server.Settings.Log.With(slog.String("state", "restarting"))

	log.Info("Restarting process")

	err := fsm.server.Settings.StorePID(-1)
	if err != nil {
		log.Error("Failed to store PID", "error", err.Error())
	}

	if s.restartCtx.LimitReached() {
		log.Error("Max restarts reached")
		fsm.handleError(errors.New("max restarts reached"))
		s.restartCtx.Reset()
		return
	}

	s.restartCtx.Increment()

	_ = fsm.proc.Stop()

	if ok, err := fsm.updater.IsNewVersionAvailable(); err == nil && ok {
		log.Info("New version available, running update")
		fsm.ChangeState(StateTUpdating)
		return
	}

	fsm.ChangeState(StateTRunning)
}

type StateUpdating struct {
	stateEmpty
}

func (s *StateUpdating) Enter(fsm *FSM) {
	_, err := fsm.updater.Update()
	if err != nil {
		fsm.handleError(err)
	}

	fsm.ChangeState(StateTRestarting)
}
