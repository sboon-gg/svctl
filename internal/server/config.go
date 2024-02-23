package server

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type ValuesSource struct {
	File string `yaml:"file"`
}

type Config struct {
	Values []ValuesSource `yaml:"values"`
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

func writeConfig(path string, conf *Config) error {
	content, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(path, ConfigFile), content, 0644)
}
