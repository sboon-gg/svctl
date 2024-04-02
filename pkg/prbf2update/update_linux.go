package prbf2update

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
)

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
	cmd.Stdout = buf
	cmd.Stderr = buf

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
