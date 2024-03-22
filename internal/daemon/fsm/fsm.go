package fsm

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"
)

var ErrActionNotAllowed = errors.New("action not allowed")

type State interface {
	Enter(*FSM)
	Exit()
}

type FSM struct {
	states         map[StateT]State
	allowedActions map[Action][]State

	currentState State
	desiredState State

	path string
	proc *os.Process

	err error

	onStateChange func(StateT)
	render        func() error

	cancel context.CancelFunc
}

func New(path string, onStateChange func(StateT), render func() error) *FSM {
	states := map[StateT]State{
		StateTStopped:    &StateStopped{},
		StateTRunning:    &StateRunning{},
		StateTRestarting: &StateRestarting{},
	}

	allowedActions := map[Action][]State{
		ActionStart: {
			states[StateTStopped],
		},
		ActionStop: {
			states[StateTRunning],
		},
		ActionRestart: {
			states[StateTRunning],
		},
		ActionAdopt: {
			states[StateTStopped],
		},
	}

	return &FSM{
		states:         states,
		allowedActions: allowedActions,
		currentState:   states[StateTStopped],
		desiredState:   states[StateTStopped],
		onStateChange:  onStateChange,
		render:         render,
		path:           path,
	}
}

func (fsm *FSM) loop() {
	if fsm.cancel != nil {
		fsm.cancel()
	}

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
	if fsm.proc != nil {
		return fsm.proc.Pid
	}

	return -1
}

func (fsm *FSM) Start() error {
	go fsm.loop()
	return fsm.action(ActionStart, StateTRunning)
}

func (fsm *FSM) Stop() error {
	err := fsm.action(ActionStop, StateTStopped)
	if err != nil {
		return err
	}

	time.Sleep(300 * time.Millisecond)
	fsm.cancel()

	return nil
}

func (fsm *FSM) Restart() error {
	return fsm.action(ActionRestart, StateTRestarting)
}

func (fsm *FSM) Adopt(proc *os.Process) error {
	if !fsm.isActionAllowed(ActionAdopt) {
		return ErrActionNotAllowed
	}

	fsm.proc = proc

	go fsm.loop()

	return fsm.Start()
}

func (fsm *FSM) isActionAllowed(action Action) bool {
	allowedStates, ok := fsm.allowedActions[action]
	if !ok {
		return false
	}

	for _, state := range allowedStates {
		if fsm.currentState == state {
			return true
		}
	}

	return false
}

func (fsm *FSM) action(a Action, s StateT) error {
	if !fsm.isActionAllowed(a) {
		return ErrActionNotAllowed
	}

	fsm.ChangeState(s)
	return nil
}

func (fsm *FSM) ChangeState(state StateT) {
	if desiredState, ok := fsm.states[state]; ok {
		fsm.desiredState = desiredState
	}
}

func (fsm *FSM) Transition() {
	if fsm.desiredState != fsm.currentState {
		slog.Debug(fmt.Sprintf("Transitioning from %T to %T", fsm.currentState, fsm.desiredState), slog.Int("pid", fsm.Pid()))
		if fsm.currentState != nil {
			fsm.currentState.Exit()
		}
		fsm.currentState = fsm.desiredState
		fsm.currentState.Enter(fsm)
		fsm.onStateChange(fsm.stateToType(fsm.currentState))
	}
}

func (fsm *FSM) stateToType(s State) StateT {
	for t, state := range fsm.states {
		if state == s {
			return t
		}
	}

	return StateTStopped
}

func (fsm *FSM) handleError(err error) {
	fsm.err = err
	slog.Error(err.Error(), slog.Int("pid", fsm.Pid()))
	fsm.ChangeState(StateTStopped)
}
