package prbf2update

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// Validating server license... Valid
//   IP: 45.63.41.122
//   Port: 16567
//   User: Vista

// Reading current version... 1.7.0.0
// Checking latest version... 1.7.4.5

const (
	exe = "prserverupdater-linux64"
)

func (u *PRBF2Update) update() ([]byte, error) {
	binPath := filepath.Join(u.path, updaterPath)
	exePath := filepath.Join(binPath, exe)

	info, err := os.Stat(exePath)
	if err != nil {
		return nil, err
	}

	if info.Mode().Perm()&0100 == 0 {
		err = os.Chmod(exePath, info.Mode().Perm()|0100)
		if err != nil {
			return nil, err
		}
	}

	cmd := exec.Command("./" + exe)
	cmd.Dir = binPath

	buf := new(bytes.Buffer)
	cmd.Stdout = io.MultiWriter(os.Stdout, buf)
	cmd.Stderr = io.MultiWriter(os.Stderr, buf)

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
