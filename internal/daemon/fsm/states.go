package fsm

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/sboon-gg/svctl/pkg/prbf2proc"
)

type stateEmpty struct{}

func (s *stateEmpty) Enter(p *FSM) {}
func (s *stateEmpty) Exit()        {}

type StateStopped struct {
	stateEmpty
}

func (s *StateStopped) Enter(fsm *FSM) {
	err := prbf2proc.Stop(fsm.proc)
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

	if fsm.proc == nil {
		err := fsm.render()
		if err != nil {
			fsm.handleError(err)
			return
		}

		proc, err := prbf2proc.Start(fsm.path)
		if err != nil {
			fsm.handleError(err)
			return
		}

		fsm.proc = proc
	}

	go func() {
		_, _ = fsm.proc.Wait()

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if !prbf2proc.IsHealthy(fsm.proc) {
					cancel()
					fsm.ChangeState(StateTRestarting)
				}

				time.Sleep(500 * time.Millisecond)
				continue
			}
		}
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
}

func (s *StateRestarting) Enter(fsm *FSM) {
	if fsm.restartCtx.max() {
		fsm.handleError(errors.New("max restarts reached"))
		fsm.restartCtx.reset()
		return
	}

	if fsm.restartCtx.count == 0 {
		err := prbf2proc.Stop(fsm.proc)
		if err != nil {
			fsm.handleError(err)
			return
		}
	}
	_ = prbf2proc.Stop(fsm.proc)

	fsm.ChangeState(StateTRunning)
}

type StateUpdating struct {
}
