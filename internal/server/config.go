package server

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Config struct {
	IP   string `yaml:"ip"`
	Port int    `yaml:"port"`
}

var DefaultConfig = Config{
	IP:   "",
	Port: 16567,
}

func (s *Server) Config() (*Config, error) {
	var config Config

	content, err := os.ReadFile(s.dotPath(ConfigFile))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func writeDefaultConfig(path string) error {
	content, err := yaml.Marshal(DefaultConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(path, ConfigFile), content, 0755)
}
