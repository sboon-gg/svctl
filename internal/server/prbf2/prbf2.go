package prbf2

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"
)

const maxRestarts = 5

type restartCtx struct {
	count       int
	resetCancel context.CancelFunc
}

func (r *restartCtx) inc() error {
	r.count++

	if r.count >= maxRestarts {
		return errors.New("max restarts reached")
	}

	if r.resetCancel != nil {
		r.resetCancel()
	}
	go r.start()

	return nil
}

func (r *restartCtx) start() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	resetCtx, resetCancel := context.WithCancel(context.Background())
	r.resetCancel = resetCancel

	for {
		select {
		case <-resetCtx.Done():
			return
		case <-ctx.Done():
			r.count = 0
			return
		}
	}
}

type PRBF2 struct {
	path          string
	onStateChange []func(State)
	// log  *slog.Logger

	proc    *os.Process
	toAdopt *os.Process

	currentState State
	desiredState State

	err error

	watchStateCancel   context.CancelFunc
	watchProcessCancel context.CancelFunc

	restartCtx restartCtx
}

func New(path string, onStateChange ...func(State)) *PRBF2 {
	return &PRBF2{
		path:          path,
		onStateChange: onStateChange,
	}
}

func (p *PRBF2) Adopt(proc *os.Process) error {
	p.toAdopt = proc
	go p.stateWatcher()
	return p.action(ActionAdopt)
}

func (p *PRBF2) Err() error {
	return p.err
}

func (p *PRBF2) Pid() int {
	if p.proc == nil {
		return -1
	}

	return p.proc.Pid
}

func (p *PRBF2) Start() error {
	go p.stateWatcher()
	return p.action(ActionStart)
}

func (p *PRBF2) Stop() error {
	err := p.action(ActionStop)
	if err != nil {
		return err
	}

	p.watchStateCancel()
	return nil
}

func (p *PRBF2) Restart() error {
	return p.action(ActionRestart)
}

var ErrNoActionTransition = errors.New("no action transition")

func (p *PRBF2) action(a Action) error {
	slog.Info("New action", slog.String("action", a.String()), slog.String("currentState", p.currentState.String()))
	transition, ok := actionTransitions[p.currentState][a]
	if !ok {
		return ErrNoActionTransition
	}

	p.currentState = transition[0]
	p.desiredState = transition[1]

	slog.Info("New states", slog.String("currentState", p.currentState.String()), slog.String("desiredState", p.desiredState.String()))

	return nil
}

func (p *PRBF2) stateWatcher() {
	watchCtx, watchCancel := context.WithCancel(context.Background())
	p.watchStateCancel = watchCancel

	for {
		select {
		case <-watchCtx.Done():
			return
		default:
			if p.currentState == p.desiredState {
				continue
			}

			transition, ok := stateTransition[p.currentState][p.desiredState]
			if !ok {
				// TODO: figure out intermediate state to reach desired state
				slog.Error(fmt.Sprintf("no state transition from %s to %s", p.currentState, p.desiredState))
				continue
			}

			transition(p)

			for _, f := range p.onStateChange {
				f(p.currentState)
			}
		}
	}
}

func (p *PRBF2) watchProcess() error {
	if p.proc == nil {
		return errors.New("no process to watch")
	}

	watchCtx, watchCancel := context.WithCancel(context.Background())
	p.watchProcessCancel = watchCancel

	go func() {
		watchCh := watchProcess(p.proc, watchCtx)
		for {
			select {
			case <-watchCtx.Done():
				return
			case _, ok := <-watchCh:
				if !ok {
					p.currentState = StateExited
					return
				}
			}
		}
	}()

	return nil
}
