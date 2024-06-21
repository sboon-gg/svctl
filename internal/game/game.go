package game

import (
	"context"
	"io"
	"os"
	"path/filepath"
)

const (
	updaterPath = "mods/pr/bin"
)

type Server struct {
	path       string
	processPID *int
}

func Open(path string) (*Server, error) {
	//TODO: adopt and store PID in directory
	return &Server{
		path: path,
	}, nil
}

func (s *Server) WriteFile(path string, data []byte) error {
	return os.WriteFile(filepath.Join(s.path, path), data, 0644)
}

func (s *Server) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(filepath.Join(s.path, path))
}

func (s *Server) Update(ctx context.Context, outW, inW, errW io.Writer) error {
	return s.update(ctx, outW, inW, errW)
}

func makeFileExecutable(exePath string) error {
	info, err := os.Stat(exePath)
	if err != nil {
		return err
	}

	if info.Mode().Perm()&0100 == 0 {
		err = os.Chmod(exePath, info.Mode()|0100)
		if err != nil {
			return err
		}
	}

	return nil
}
