package fsm

import (
	"context"
	"time"
)

type StateRunning struct {
	baseState
	counter *restartCounter
	cancel  context.CancelFunc
}

func NewStateRunning(counter *restartCounter) *StateRunning {
	if counter == nil {
		counter = &restartCounter{}
	}

	return &StateRunning{
		counter: counter,
	}
}

func (s *StateRunning) OnEnter(fsm *FSM) {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	ticker := time.NewTicker(time.Minute)

	sv := fsm.Server()

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				fsm.Server().Render()
			default:
				if !sv.IsRunning() {
					fsm.ChangeState(NewStateRestarting(s.counter))
					ticker.Stop()
					cancel()
					return
				}

				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
}

func (s *StateRunning) OnExit() {
	s.cancel()
}

func (s *StateRunning) EventHandler(event Event, fsm *FSM) (State, error) {
	switch event {
	case EventStop:
		if err := fsm.Server().Stop(); err != nil {
			return NewStateErrored(err), err
		}

		return NewStateStopped(), nil
	case EventRestart:
		if err := fsm.Server().Stop(); err != nil {
			return NewStateErrored(err), err
		}

		return NewStateRestarting(s.counter), nil
	default:
		return nil, ErrEventNotAllowed
	}
}
