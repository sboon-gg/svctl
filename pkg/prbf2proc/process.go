package prbf2proc

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type PRBF2Process struct {
	path string

	adopted bool
	process *os.Process
}

func New(path string) (*PRBF2Process, error) {
	err := verifyPath(path)
	if err != nil {
		return nil, fmt.Errorf("Path %q is not a PRBF2 server", path)
	}

	return &PRBF2Process{
		path: path,
	}, nil
}

func (p *PRBF2Process) Adopt(proc *os.Process) error {
	p.process = proc

	if !p.IsRunning() {
		p.process = nil
		return fmt.Errorf("Process %d is not running", proc.Pid)
	}

	p.adopted = true

	return nil
}

func (p *PRBF2Process) Pid() int {
	if p.process == nil {
		return -1
	}

	return p.process.Pid
}

func (p *PRBF2Process) Start() error {
	if p.process != nil {
		return nil
	}

	proc, err := startProcess(p.path)
	if err != nil {
		return err
	}

	p.process = proc

	return nil
}

func (p *PRBF2Process) Stop() error {
	if p.process == nil {
		return nil
	}

	err := p.process.Kill()
	if err != nil {
		return err
	}

	_ = p.process.Release()

	p.process = nil
	p.adopted = false

	return nil
}

func (p *PRBF2Process) IsRunning() bool {
	if p.process == nil {
		return false
	}

	proc, err := process.NewProcess(int32(p.process.Pid))
	if err != nil {
		return false
	}

	isRunning, err := proc.IsRunning()
	if err != nil {
		return false
	}

	return isRunning
}

func (p *PRBF2Process) Wait() error {
	if p.adopted {
		p.waitForAdopted()
		return nil
	}

	_, err := p.process.Wait()
	return err
}

func (p *PRBF2Process) waitForAdopted() {
	for {
		if !p.IsRunning() {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func verifyPath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	_, err = os.Stat(filepath.Join(path, "mods/pr/mod.desc"))
	return err
}
