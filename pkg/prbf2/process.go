package prbf2

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Status int

const (
	StatusStopped Status = iota // StatusStopped
	StatusRunning               // StatusRunning
	StatusExited                // StatusExited
)

const (
	updaterPath = "mods/pr/bin"
)

var ErrAlreadyRunning = errors.New("process already running")
var ErrNotRunning = errors.New("process not running")

type Process struct {
	path string

	proc   *os.Process
	status Status

	watchCancel context.CancelFunc
}

func New(path string) *Process {
	return &Process{
		path: path,
	}
}

func (p *Process) Status() Status {
	return p.status
}

func (p *Process) Adopt(proc *os.Process) error {
	if p.proc != nil {
		return ErrAlreadyRunning
	}

	p.proc = proc
	p.status = StatusRunning

	return p.watchProcess()
}

func (p *Process) Update() error {
	if p.status == StatusRunning {
		return errors.New("Game process is running")
	}

	path := filepath.Join(p.path, "mods/pr/bin")
	if runtime.GOOS == "windows" {
		path = filepath.Join(path, "prserverupdater-win32.exe")
	} else {
		path = filepath.Join(path, "prserverupdater-linux64")
	}

	cmd := exec.Command(path)
	return cmd.Run()
}

func (p *Process) Pid() int {
	if p.proc == nil {
		return -1
	}

	return p.proc.Pid
}

func (p *Process) Start() error {
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

func (p *Process) Stop() error {
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

func (p *Process) watchProcess() error {
	if p.proc == nil {
		return errors.New("no process to watch")
	}

	watchCtx, watchCancel := context.WithCancel(context.Background())
	p.watchCancel = watchCancel

	go func() {
		watchCh := watchProcess(watchCtx, p.proc)
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

func Update(path string) ([]string, error) {
	_, err := update(path)
	if err != nil {
		return nil, err
	}

	// TODO: changed files
	var changedFiles []string

	return changedFiles, nil
}
