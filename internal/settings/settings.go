package settings

import (
	"os"
	"path/filepath"

	"github.com/sboon-gg/svctl/pkg/templates"
)

const (
	SvctlDir     = ".svctl"
	TemplatesDir = "templates"

	ConfigFile        = "config.yaml"
	defaultValuesFile = "values.yaml"

	CacheFile = ".cache.yaml"
)

type Settings struct {
	path      string
	Templates *templates.Renderer
}

func Open(path string) (*Settings, error) {
	// TODO: validate settings
	s := &Settings{
		path: path,
	}

	_, err := s.Config()
	if err != nil {
		return nil, err
	}

	templatesPath := filepath.Join(path, TemplatesDir)
	_, err = os.Stat(templatesPath)
	if err == nil {
		t, err := templates.NewFromPath(templatesPath)
		if err != nil {
			return nil, err
		}

		s.Templates = t
	}

	return s, nil
}

type Opts struct {
	TemplatesRepo string
	Token         string
}

func Initialize(path string, opts *Opts) (*Settings, error) {
	if opts == nil {
		opts = &Opts{}
	}

	err := os.Mkdir(path, 0755)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if opts.TemplatesRepo != "" {
		templatesPath := filepath.Join(path, TemplatesDir)
		err = cloneTemplates(templatesPath, opts.TemplatesRepo, opts.Token)
		if err != nil {
			return nil, err
		}

		t, err := templates.NewFromPath(templatesPath)
		if err != nil {
			return nil, err
		}

		defaultsContent, err := t.DefaultsContent()
		if err != nil {
			return nil, err
		}

		err = writeValues(path, defaultsContent)
		if err != nil {
			return nil, err
		}

		config.Values = append(config.Values, ValuesSource{
			File: defaultValuesFile,
		})
	}

	err = writeConfig(path, config)
	if err != nil {
		return nil, err
	}

	return Open(path)
}
