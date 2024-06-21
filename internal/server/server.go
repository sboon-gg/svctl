package server

import (
	"github.com/sboon-gg/svctl/internal/game"
	"github.com/sboon-gg/svctl/internal/settings"
	"github.com/sboon-gg/svctl/pkg/templates"
)

type Server struct {
	game.Server
	settings.Settings
}

func Open(serverPath, settingsPath string) (*Server, error) {
	s, err := settings.Open(settingsPath)
	if err != nil {
		return nil, err
	}

	g, err := game.Open(serverPath)
	if err != nil {
		return nil, err
	}

	return &Server{
		Server:   *g,
		Settings: *s,
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

	outputs, err := s.Settings.Templates.Render(values)
	if err != nil {
		return err
	}

	for _, output := range outputs {
		err = s.WriteFile(output.Destination, output.Content)
		if err != nil {
			return err
		}
	}

	return nil
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
