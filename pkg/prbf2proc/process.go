package prbf2proc

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

func Start(path string) (*os.Process, error) {
	err := verifyPath(path)
	if err != nil {
		return nil, fmt.Errorf("Path %q is not a PRBF2 server", path)
	}

	return startProcess(path)
}

func Stop(process *os.Process) error {
	err := process.Kill()
	if err != nil {
		return err
	}

	_ = process.Release()
	return nil
}

func Wait(process *os.Process) context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		_, err := process.Wait()
		if err == nil {
			_ = process.Release()
		}

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if isProcessHealthy(process) {
					time.Sleep(500 * time.Millisecond)
					continue
				}

				cancel()
				return
			}
		}
	}()

	return ctx
}

func isProcessHealthy(process *os.Process) bool {
	err := process.Signal(syscall.Signal(0))
	return err == nil
}

func verifyPath(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return err
	}

	_, err = os.Stat(filepath.Join(path, "mods/pr/mod.desc"))
	return err
}
