package fsm

import (
	"context"
	"time"

	"github.com/sboon-gg/svctl/internal/server"
)

type FSM struct {
	server *server.Server

	currentState State
	desiredState State

	cancel context.CancelFunc
}

func New(server *server.Server, initialState State) *FSM {
	return &FSM{
		currentState: initialState,
		server:       server,
	}
}

func (f *FSM) Server() *server.Server {
	return f.server
}

func (f *FSM) ChangeState(state State) {
	f.desiredState = state
}

func (f *FSM) Event(event Event) error {
	if f.currentState == nil {
		return nil
	}

	// Error is for the user, state is for the FSM
	nextState, err := f.currentState.EventHandler(event, f)
	if nextState != nil {
		f.desiredState = nextState
	}

	if err != nil {
		return err
	}

	return nil
}

func (f *FSM) Run() {
	if f.cancel != nil {
		f.cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	f.cancel = cancel

	timer := time.NewTimer(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			f.Transition()
		}
	}
}

func (f *FSM) Transition() {
	if f.desiredState != f.currentState {
		if f.currentState != nil {
			f.currentState.OnExit()
		}

		f.currentState = f.desiredState
		f.currentState.OnEnter(f)
	}
}
