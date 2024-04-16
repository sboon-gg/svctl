package fsm

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/sboon-gg/svctl/internal/server"
	"github.com/sboon-gg/svctl/pkg/prbf2proc"
	"github.com/sboon-gg/svctl/pkg/prbf2update"
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

	server *server.Server

	proc    *prbf2proc.PRBF2Process
	updater *prbf2update.PRBF2Update

	err error

	cancel context.CancelFunc
}

func New(sv *server.Server, updateCache *prbf2update.Cache) *FSM {
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

	// Ignore error since we know the path is valid
	proc, _ := prbf2proc.New(sv.Path)

	return &FSM{
		states:         states,
		allowedActions: allowedActions,
		currentState:   states[StateTStopped],
		desiredState:   states[StateTStopped],
		server:         sv,
		proc:           proc,
		updater:        prbf2update.New(sv.Path, updateCache),
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
	return fsm.proc.Pid()
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

	err := fsm.proc.Adopt(proc)
	if err != nil {
		return err
	}

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
		fsm.server.Settings.Log.Debug(fmt.Sprintf("Transitioning from %T to %T", fsm.currentState, fsm.desiredState))
		if fsm.currentState != nil {
			fsm.currentState.Exit()
		}
		fsm.currentState = fsm.desiredState
		fsm.currentState.Enter(fsm)
	}
}

func (fsm *FSM) handleError(err error) {
	fsm.err = err
	fsm.server.Settings.Log.Error(err.Error())
	fsm.ChangeState(StateTStopped)
}
