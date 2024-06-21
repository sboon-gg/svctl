package game

import (
	"errors"
	"runtime"

	"github.com/shirou/gopsutil/v3/process"
)

// Should the process be zeroed if it stops by itself and IsRunning() returns false?

var (
	ErrProcessAlreadyRunning = errors.New("process already running")
)

var commonProcessArgs = []string{
	"+modPath", "mods/pr",
	"+noStatusMonitor", "1",
	"+multi", "1",
	"+dedicated", "1",
}

func (s *Server) Start() error {
	if s.IsRunning() {
		return ErrProcessAlreadyRunning
	}

	proc, err := s.startProcess()
	if err != nil {
		return err
	}

	pid := proc.Pid
	s.processPID = &pid

	return nil
}

func (s *Server) Stop() error {
	if !s.IsRunning() {
		return nil
	}

	proc, err := process.NewProcess(int32(*s.processPID))
	if err != nil {
		return err
	}

	err = proc.Kill()
	if err != nil {
		return err
	}

	s.clearProcessPID()

	return nil
}

func (s *Server) IsRunning() bool {
	if s.processPID == nil {
		return false
	}

	// On Windows we need to check if the process isn't hanging on an error dialog.
	if runtime.GOOS == "windows" {
		health, err := processHealth(*s.processPID)
		if err == nil && !health {
			s.clearProcessPID()
			return false
		}
	}

	proc, err := process.NewProcess(int32(*s.processPID))
	if err != nil {
		return false
	}

	exe, err := proc.Exe()
	if err != nil {
		return false
	}

	if s.processExe() != exe {
		// Not our process.
		s.clearProcessPID()
		return false
	}

	isRunning, err := proc.IsRunning()
	if err != nil {
		return false
	}

	return isRunning
}

func (s *Server) clearProcessPID() {
	s.processPID = nil
}
