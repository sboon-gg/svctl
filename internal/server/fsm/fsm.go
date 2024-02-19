package fsm

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/sboon-gg/svctl/internal/server/fsm/prbf2"
)

var ErrActionNotAllowed = errors.New("action not allowed")

type State interface {
	Enter(*FSM)
	Type() StateT
	Exit()
}

var allowedActions = map[Action][]StateT{
	ActionStart: {
		StateTStopped,
	},
	ActionStop: {
		StateTRunning,
		StateTErrored,
	},
	ActionRestart: {
		StateTRunning,
	},
	ActionAdopt: {
		StateTStopped,
	},
}

type FSM struct {
	currentState State
	desiredState State

	ctrl *prbf2.PRBF2

	restartCtx restartCtx
	err        error

	onStateChange func(StateT)

	cancel context.CancelFunc
}

func New(path string, onStateChange func(StateT)) *FSM {
	return &FSM{
		ctrl:          prbf2.New(path),
		currentState:  &StateStopped{},
		desiredState:  &StateStopped{},
		onStateChange: onStateChange,
	}
}

func (fsm *FSM) loop() {
	ctx, cancel := context.WithCancel(context.Background())
	fsm.cancel = cancel

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fsm.Transition()
		}
	}
}

func (fsm *FSM) Pid() int {
	return fsm.ctrl.Pid()
}

func (fsm *FSM) Start() error {
	go fsm.loop()
	return fsm.action(ActionStart, &StateRunning{})
}

func (fsm *FSM) Stop() error {
	err := fsm.action(ActionStop, &StateStopped{})
	if err != nil {
		return err
	}

	time.Sleep(300 * time.Millisecond)
	fsm.cancel()

	return nil
}

func (fsm *FSM) Restart() error {
	return fsm.action(ActionRestart, &StateRestarting{})
}

func (fsm *FSM) Adopt(proc *os.Process) error {

	if !fsm.isActionAllowed(ActionAdopt) {
		return ErrActionNotAllowed
	}

	err := fsm.ctrl.Adopt(proc)
	if err != nil {
		return err
	}

	go fsm.loop()

	fsm.ChangeState(&StateRunning{})
	return nil
}

func (fsm *FSM) isActionAllowed(action Action) bool {
	allowedStates, ok := allowedActions[action]
	if !ok {
		return false
	}

	for _, state := range allowedStates {
		if fsm.currentState.Type() == state {
			return true
		}
	}

	return false
}

func (fsm *FSM) action(a Action, s State) error {
	if !fsm.isActionAllowed(a) {
		return ErrActionNotAllowed
	}

	fsm.ChangeState(s)
	return nil
}

func (fsm *FSM) ChangeState(state State) {
	fsm.desiredState = state
}

func (fsm *FSM) Transition() {
	if fsm.desiredState != fsm.currentState {
		if fsm.currentState != nil {
			fsm.currentState.Exit()
		}
		fsm.currentState = fsm.desiredState
		fsm.currentState.Enter(fsm)
		fsm.onStateChange(fsm.currentState.Type())
	}
}

func (fsm *FSM) handleError(err error) {
	fsm.err = err
	log.Print(err.Error())
	fsm.ChangeState(&StateStopped{})
}
