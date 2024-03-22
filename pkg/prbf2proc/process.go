package prbf2proc

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func Start(path string) (*os.Process, error) {
	err := verifyPath(path)
	if err != nil {
		return nil, fmt.Errorf("Path %q is not a PRBF2 server", path)
	}

	return startProcess(path)
}

func Stop(process *os.Process) error {
	if process == nil {
		return nil
	}

	err := process.Kill()
	if err != nil {
		return err
	}

	_ = process.Release()
	return nil
}

func IsHealthy(process *os.Process) bool {
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
