package server

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sboon-gg/svctl/pkg/templates"
)

const (
	SvctlDir     = ".svctl"
	TemplatesDir = "templates"

	ConfigFile        = "config.yaml"
	defaultValuesFile = "values.yaml"

	CacheFile = ".cache.yaml"
)

type Server struct {
	Path string
}

func Open(serverPath string) (*Server, error) {
	_, err := os.Stat(filepath.Join(serverPath, SvctlDir))
	if err != nil {
		return nil, errors.Wrap(err, "server not initialized")
	}

	return &Server{
		Path: serverPath,
	}, nil
}

type Opts struct {
	TemplatesRepo string
	Token         string
}

func Initialize(serverPath string, opts *Opts) (*Server, error) {
	if opts == nil {
		opts = &Opts{}
	}

	svctlPath := filepath.Join(serverPath, SvctlDir)
	err := os.Mkdir(svctlPath, 0755)
	if err != nil {
		return nil, err
	}

	config := &Config{}

	if opts.TemplatesRepo != "" {
		templatesPath := filepath.Join(svctlPath, TemplatesDir)
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

		err = writeValues(svctlPath, defaultsContent)
		if err != nil {
			return nil, err
		}

		config.Values = append(config.Values, ValuesSource{
			File: defaultValuesFile,
		})
	}

	err = writeConfig(svctlPath, config)
	if err != nil {
		return nil, err
	}

	return Open(serverPath)
}
