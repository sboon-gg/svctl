package server

import (
	"os"
	"path/filepath"

	"github.com/sboon-gg/svctl/pkg/templates"
)

func (s *Server) Render() ([]templates.RenderOutput, error) {
	if s.Settings.Templates == nil {
		return []templates.RenderOutput{}, nil
	}

	values, err := s.Settings.Values()
	if err != nil {
		return nil, err
	}

	return s.Settings.Templates.Render(values)
}

func (s *Server) WriteTemplatesOutput(outputs []templates.RenderOutput) error {
	for _, out := range outputs {
		path := filepath.Join(s.ServerPath, out.Destination)

		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(path, out.Content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
