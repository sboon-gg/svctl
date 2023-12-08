package cmd

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/sboon-gg/prbf2-templates/pkg/config"
	"github.com/sboon-gg/prbf2-templates/pkg/templates"
	"github.com/sboon-gg/prbf2-templates/pkg/values"
	"gopkg.in/yaml.v3"
)

const (
	dotDir = ".templater"
)

type serverInstance struct {
	path string
}

func newServerInstance(path string) (*serverInstance, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	return &serverInstance{
		path: path,
	}, nil
}

func (si *serverInstance) dotDir() string {
	return filepath.Join(si.path, dotDir)
}

func (si *serverInstance) HasConfigCache() bool {
	if _, err := os.Stat(si.dotDir()); os.IsNotExist(err) {
		return false
	}
	return true
}

func (si *serverInstance) CreateConfigDir() error {
	return os.Mkdir(si.dotDir(), os.ModePerm)
}

func (si *serverInstance) WriteDefaultConfig() error {
	if !si.HasConfigCache() {
		if err := si.CreateConfigDir(); err != nil {
			return err
		}
	}

	out, err := yaml.Marshal(config.DefaultConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(si.dotDir(), "config.yaml"), out, 0755)
}

func (si *serverInstance) WriteDefaultTemplates() error {
	if err := os.MkdirAll(filepath.Join(si.dotDir(), "templates"), os.ModePerm); err != nil {
		return err
	}

	return fs.WalkDir(templates.DefaultTemplateFiles, "templates", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		content, err := fs.ReadFile(templates.DefaultTemplateFiles, path)
		if err != nil {
			return err
		}

		return os.WriteFile(filepath.Join(si.dotDir(), path), content, 0755)
	})
}

func (si *serverInstance) WriteDefaultValues() error {
	if !si.HasConfigCache() {
		if err := si.CreateConfigDir(); err != nil {
			return err
		}
	}

	return os.WriteFile(filepath.Join(si.dotDir(), "values.yaml"), values.DefaultValuesFile, 0755)
}
