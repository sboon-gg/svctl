package prbf2proc

import (
	"context"
	"errors"
	"os"
	"time"
)

type Status int

const (
	StatusStopped Status = iota // StatusStopped
	StatusRunning               // StatusRunning
	StatusExited                // StatusExited
)

var ErrAlreadyRunning = errors.New("process already running")
var ErrNotRunning = errors.New("process not running")

type GameProcess struct {
	path string

	process *os.Process
	status  Status

	cancel context.CancelFunc
}

func NewGameProcess(path string) *GameProcess {
	return &GameProcess{path: path}
}

func (gs *GameProcess) Status() Status {
	return gs.status
}

func (gs *GameProcess) Pid() int {
	if gs.process == nil {
		return -1
	}

	return gs.process.Pid
}

func (gs *GameProcess) Adopt(proc *os.Process) error {
	if gs.process != nil {
		return ErrAlreadyRunning
	}

	gs.process = proc
	gs.status = StatusRunning

	return nil
}

func (gs *GameProcess) Start() error {
	if gs.process != nil {
		return ErrAlreadyRunning
	}

	proc, err := Start(gs.path)
	if err != nil {
		return err
	}

	gs.process = proc
	gs.status = StatusRunning

	return gs.watch()
}

func (gs *GameProcess) Stop() error {
	if gs.process == nil {
		return ErrNotRunning
	}

	gs.cancel()

	err := Stop(gs.process)
	if err != nil {
		return err
	}

	gs.status = StatusStopped
	gs.process = nil

	return nil
}

func (gs *GameProcess) watch() error {
	if gs.process == nil {
		return ErrNotRunning
	}

	ctx, cancel := context.WithCancel(context.Background())
	gs.cancel = cancel

	go func() {
		waitCtx := Wait(gs.process)
		for {
			select {
			case <-waitCtx.Done():
				if gs.status == StatusRunning {
					gs.status = StatusExited
					gs.process = nil
				}
				return
			case <-ctx.Done():
				return
			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	return nil
}
