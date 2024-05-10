package fsm

import (
	"context"
	"time"

	"github.com/sboon-gg/svctl/internal/game"
)

type Server interface {
	game.Game
	Render() error
}

type FSM interface {
	Server() Server
	ChangeState(State)
	Event(Event) error
}

type fsm struct {
	server Server

	currentState State
	desiredState State

	cancel context.CancelFunc
}

func New(server Server, initialState State) FSM {
	return &fsm{
		currentState: initialState,
		server:       server,
	}
}

func (f *fsm) Server() Server {
	return f.server
}

func (f *fsm) ChangeState(state State) {
	f.desiredState = state
}

func (f *fsm) Event(event Event) error {
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

func (f *fsm) Run() {
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

func (f *fsm) Transition() {
	if f.desiredState != f.currentState {
		if f.currentState != nil {
			f.currentState.OnExit()
		}

		f.currentState = f.desiredState
		f.currentState.OnEnter(f)
	}
}
