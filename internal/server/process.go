package server

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"syscall"
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

func isHealthy(process *os.Process) bool {
	err := process.Signal(syscall.Signal(0))
	return err == nil
}

func stopProcess(process *os.Process) error {
	if process == nil || process.Pid == -1 || !isHealthy(process) {
		return nil
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		_, _ = process.Wait()
	}(wg)

	if isHealthy(process) {
		err := process.Kill()
		if err != nil {
			return err
		}
	}

	wg.Wait()

	return process.Release()
}
