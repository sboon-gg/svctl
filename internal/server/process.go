package server

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

const (
	exe       = "prbf2_l64ded"
	binaryDir = "bin/amd-64"
)

func startProcess(path string) (*os.Process, error) {
	binDir := filepath.Join(path, binaryDir)
	fullExe := filepath.Join(binDir, exe)

	env := os.Environ()
	env = append(env, fmt.Sprintf("LD_LIBRARY_PATH=%s", binDir))

	args := []string{
		fullExe,
		"+modPath", "mods/pr",
		"+noStatusMonitor", "1",
		"+multi", "1",
		"+dedicated", "1",
	}

	return os.StartProcess(fullExe, args, &os.ProcAttr{
		Dir: path,
		Env: env,
		Sys: &syscall.SysProcAttr{
			Setpgid: true,
		},
	})
}

func healthCheck(process *os.Process) error {
	if process == nil {
		return nil
	}

	return process.Signal(syscall.Signal(0))
}

func stopProcess(process *os.Process) error {
	if process == nil {
		return nil
	}

	if process.Pid == -1 {
		return nil
	}

	calls := []syscall.Signal{
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	}

	for _, call := range calls {
		if err := healthCheck(process); err != nil {
			return nil
		}

		// send signal to process
		if err := process.Signal(call); err != nil {
			return err
		}

		// give process time to shutdown
		time.Sleep(500 * time.Millisecond)
	}

	err := process.Kill()
	if err != nil {
		return err
	}

	return process.Release()
}
