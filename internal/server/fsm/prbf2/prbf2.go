package prbf2

import (
	"context"
	"errors"
	"os"
)

type Status int

const (
	StatusStopped Status = iota // StatusStopped
	StatusRunning               // StatusRunning
	StatusExited                // StatusExited
)

var ErrAlreadyRunning = errors.New("process already running")
var ErrNotRunning = errors.New("process not running")

type PRBF2 struct {
	path string

	proc   *os.Process
	status Status

	watchCancel context.CancelFunc
}

func New(path string) *PRBF2 {
	return &PRBF2{
		path: path,
	}
}

func (p *PRBF2) Status() Status {
	return p.status
}

func (p *PRBF2) Adopt(proc *os.Process) error {
	if p.proc != nil {
		return ErrAlreadyRunning
	}

	p.proc = proc
	p.status = StatusRunning

	return p.watchProcess()
}

func (p *PRBF2) Pid() int {
	if p.proc == nil {
		return -1
	}

	return p.proc.Pid
}

func (p *PRBF2) Start() error {
	if p.proc != nil {
		return ErrAlreadyRunning
	}

	proc, err := startProcess(p.path)
	if err != nil {
		return err
	}

	p.proc = proc
	p.status = StatusRunning

	return p.watchProcess()
}

func (p *PRBF2) Stop() error {
	if p.proc == nil {
		return ErrNotRunning
	}

	err := stopProcess(p.proc)
	if err != nil {
		return err
	}

	p.watchCancel()
	p.status = StatusStopped
	p.proc = nil
	return nil
}

func (p *PRBF2) watchProcess() error {
	if p.proc == nil {
		return errors.New("no process to watch")
	}

	watchCtx, watchCancel := context.WithCancel(context.Background())
	p.watchCancel = watchCancel

	go func() {
		watchCh := watchProcess(p.proc, watchCtx)
		for {
			select {
			case <-watchCtx.Done():
				return
			case _, ok := <-watchCh:
				if !ok {
					p.status = StatusExited
					p.watchCancel()
					p.proc = nil
					return
				}
			}
		}
	}()

	return nil
}
