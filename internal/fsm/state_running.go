package fsm

import (
	"context"
	"errors"
	"os"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type StateRunning struct {
	baseState
	process *os.Process
	counter *restartCounter
	cancel  context.CancelFunc
}

func NewStateRunning(process *os.Process, counter *restartCounter) *StateRunning {
	if counter == nil {
		counter = &restartCounter{}
	}

	return &StateRunning{
		process: process,
		counter: counter,
	}
}

func (s *StateRunning) OnEnter(fsm FSM) {
	ctx, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	isHealthy := s.isHealthyCheck()

	ticker := time.NewTicker(time.Minute)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				fsm.Server().Render()
			default:
				if !isHealthy() {
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

func (s *StateRunning) EventHandler(event Event, fsm FSM) (State, error) {
	switch event {
	case EventStop:
		if err := s.stopProcess(); err != nil {
			return NewStateErrored(err), err
		}

		return NewStateStopped(), nil
	case EventRestart:
		if err := s.stopProcess(); err != nil {
			return NewStateErrored(err), err
		}

		return NewStateRestarting(s.counter), nil
	default:
		return nil, ErrEventNotAllowed
	}
}

func (s *StateRunning) stopProcess() error {
	err := s.process.Kill()
	if err != nil {
		return err
	}

	return s.process.Release()
}

func (s *StateRunning) isHealthyCheck() func() bool {
	proc, err := process.NewProcess(int32(s.process.Pid))
	if errors.Is(err, process.ErrorProcessNotRunning) {
		return func() bool { return false }
	} else if err != nil {
		return func() bool {
			err := s.process.Signal(syscall.Signal(0))
			return err == nil
		}
	}

	return func() bool {
		running, err := proc.IsRunning()
		if err != nil {
			panic(err)
		}
		return running
	}
}
