package server

import (
	"github.com/sboon-gg/svctl/internal/settings"
)

type Server struct {
	ServerPath string
	Settings   *settings.Settings
}

func Open(serverPath, settingsPath string) (*Server, error) {
	s, err := settings.Open(settingsPath)
	if err != nil {
		return nil, err
	}

	return &Server{
		ServerPath: serverPath,
		Settings:   s,
	}, nil
}
