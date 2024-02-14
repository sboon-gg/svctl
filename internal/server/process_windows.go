//go:build windows

package server

import (
	"os"
)

const (
	exe = "prbf2_w32ded.exe"
)

func startProcessOS(path string) (*os.Process, error) {
	allArgs := append([]string{exe}, args...)

	return os.StartProcess(exe, allArgs, &os.ProcAttr{
		Dir: path,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	})
}
