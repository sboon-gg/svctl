package server

import (
	"github.com/sboon-gg/svctl/internal/settings"
	"github.com/sboon-gg/svctl/pkg/templates"
)

type Server struct {
	Path     string
	Settings *settings.Settings
}

func Open(serverPath, settingsPath string) (*Server, error) {
	s, err := settings.Open(settingsPath)
	if err != nil {
		return nil, err
	}

	return &Server{
		Path:     serverPath,
		Settings: s,
	}, nil
}

func (s *Server) Render() error {
	if s.Settings.Templates == nil {
		return nil
	}

	values, err := s.Settings.Values()
	if err != nil {
		return err
	}

	return s.Settings.Templates.RenderInto(s.Path, values)
}

func (s *Server) DryRender() ([]templates.RenderOutput, error) {
	if s.Settings.Templates == nil {
		return nil, nil
	}

	values, err := s.Settings.Values()
	if err != nil {
		return nil, err
	}

	return s.Settings.Templates.Render(values)
}
