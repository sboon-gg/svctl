package fsm

import (
	"context"
	"errors"

	"github.com/sboon-gg/svctl/internal/server/fsm/prbf2"
)

type StateEmpty struct{}

func (s *StateEmpty) Enter(p *FSM) {}
func (s *StateEmpty) Exit()        {}

type StateStopped struct {
	StateEmpty
}

func (s *StateStopped) Type() StateT {
	return StateTStopped
}

func (s *StateStopped) Enter(fsm *FSM) {
	err := fsm.ctrl.Stop()
	if err != nil && !errors.Is(err, prbf2.ErrNotRunning) {
		fsm.handleError(err)
		return
	}
}

type StateRunning struct {
	cancel context.CancelFunc
}

func (s *StateRunning) Type() StateT {
	return StateTRunning
}

func (s *StateRunning) Enter(fsm *FSM) {
	err := fsm.ctrl.Start()
	if err != nil && !errors.Is(err, prbf2.ErrAlreadyRunning) {
		fsm.handleError(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				s.checkStatus(fsm)
			}
		}
	}()
}

func (s *StateRunning) checkStatus(fsm *FSM) {
	switch fsm.ctrl.Status() {
	case prbf2.StatusRunning:
		return
	case prbf2.StatusExited:
		fsm.ChangeState(&StateRestarting{
			onError: true,
		})
		s.cancel()
	case prbf2.StatusStopped:
		fsm.ChangeState(&StateStopped{})
		s.cancel()
	}
}

func (s *StateRunning) Exit() {
	s.cancel()
}

type StateRestarting struct {
	StateEmpty
	onError bool
}

func (s *StateRestarting) Type() StateT {
	return StateTRestarting
}

func (s *StateRestarting) Enter(fsm *FSM) {
	if s.onError {
		err := fsm.restartCtx.inc()
		if err != nil {
			fsm.handleError(err)
			return
		}
	} else {
		err := fsm.ctrl.Stop()
		if err != nil {
			fsm.handleError(err)
			return
		}
	}

	fsm.ChangeState(&StateRunning{})
}
