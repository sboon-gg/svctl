//go:build linux

package prbf2

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
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
