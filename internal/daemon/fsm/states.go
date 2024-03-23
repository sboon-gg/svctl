package fsm

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"
)

type stateEmpty struct{}

func (s *stateEmpty) Enter(p *FSM) {}
func (s *stateEmpty) Exit()        {}

type StateStopped struct {
	stateEmpty
}

func (s *StateStopped) Enter(fsm *FSM) {
	err := fsm.proc.Stop()
	if err != nil {
		fsm.handleError(err)
	}
	fsm.cancel()
}

type StateRunning struct {
	cancel context.CancelFunc
}

func (s *StateRunning) Enter(fsm *FSM) {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	if !fsm.proc.IsRunning() {
		err := fsm.render()
		if err != nil {
			fsm.handleError(err)
			return
		}

		err = fsm.proc.Start()
		if err != nil {
			fsm.handleError(err)
			return
		}
	}

	go func() {
		err := fsm.proc.Wait()
		if err != nil {
			fsm.handleError(errors.Join(err, fmt.Errorf("process exited with error")))
		}

		cancel()
		fsm.ChangeState(StateTRestarting)
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := fsm.render()
				if err != nil {
					slog.Error("Failed to render templates: %s", err, slog.Int("pid", fsm.Pid()))
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
	if s.restartCtx.LimitReached() {
		fsm.handleError(errors.New("max restarts reached"))
		s.restartCtx.Reset()
		return
	}

	s.restartCtx.Increment()

	_ = fsm.proc.Stop()

	fsm.ChangeState(StateTRunning)
}

type StateUpdating struct {
}
