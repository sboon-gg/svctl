//go:build linux

package prbf2

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"
)

const (
	exe       = "prbf2_l64ded"
	binaryDir = "bin/amd-64"
)

var args = []string{
	"+modPath", "mods/pr",
	"+noStatusMonitor", "1",
	"+multi", "1",
	"+dedicated", "1",
}

func startProcess(path string) (*os.Process, error) {
	binDir := filepath.Join(path, binaryDir)
	fullExe := filepath.Join(binDir, exe)

	env := os.Environ()
	env = append(env, fmt.Sprintf("LD_LIBRARY_PATH=%s", binDir))

	allArgs := append([]string{fullExe}, args...)

	return os.StartProcess(fullExe, allArgs, &os.ProcAttr{
		Dir: path,
		Env: env,
		Sys: &syscall.SysProcAttr{
			Setpgid: true,
		},
	})
}

func stopProcess(process *os.Process) error {
	err := process.Kill()
	if err != nil {
		return err
	}

	_ = process.Release()
	return nil
}

func update(path string) ([]byte, error) {
	path = filepath.Join(path, updaterPath, "prserverupdater-linux64")

	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if info.Mode().Perm()&0100 == 0 {
		err = os.Chmod(path, info.Mode().Perm()|0100)
		if err != nil {
			return nil, err
		}
	}

	cmd := exec.Command(path)
	return cmd.CombinedOutput()
}

func watchProcess(ctx context.Context, p *os.Process) <-chan struct{} {
	ch := make(chan struct{}, 1)

	go func() {
		_, err := p.Wait()
		if err == nil {
			_ = p.Release()
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if isProcessHealthy(p) {
					time.Sleep(500 * time.Millisecond)
					continue
				}

				ch <- struct{}{}
				close(ch)
				return
			}
		}
	}()

	return ch
}

func isProcessHealthy(process *os.Process) bool {
	err := process.Signal(syscall.Signal(0))
	return err == nil
}
