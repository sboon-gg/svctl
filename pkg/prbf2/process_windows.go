//go:build windows

package prbf2

import (
	"os"

	"golang.org/x/sys/windows"
)

const (
	exe = "prbf2_w32ded.exe"
)

var args = []string{
	"+modPath", "mods/pr",
	"+noStatusMonitor", "1",
	"+multi", "1",
	"+dedicated", "1",
}

func startProcess(path string) (*os.Process, error) {
	allArgs := append([]string{exe}, args...)

	proc, err := os.StartProcess(exe, allArgs, &os.ProcAttr{
		Dir: path,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	})
	if err != nil {
		return nil, err
	}

	setHighPriority(proc.Pid)

	return proc, nil
}

const PROCESS_ALL_ACCESS = windows.STANDARD_RIGHTS_REQUIRED | windows.SYNCHRONIZE | 0xffff

func setHighPriority(pid int) error {
	handle, err := windows.OpenProcess(PROCESS_ALL_ACCESS, false, uint32(pid))
	if err != nil {
		return err
	}

	err = windows.SetPriorityClass(handle, windows.HIGH_PRIORITY_CLASS)
	if err != nil {
		return err
	}

	return windows.CloseHandle(handle)
}
