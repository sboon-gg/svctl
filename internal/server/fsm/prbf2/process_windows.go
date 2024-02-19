//go:build windows

package prbf2

import (
	"context"
	"os"
	"syscall"
	"time"

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

	return proc, nil
}

func stopProcess(process *os.Process) error {
	err := process.Kill()
	if err != nil {
		return err
	}

	_ = process.Release()
	return nil
}

func watchProcess(ctx context.Context, p *os.Process) <-chan struct{} {
	ch := make(chan struct{}, 1)

	setHighPriority(p.Pid)

	startErrorKiller(ctx, p)

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
