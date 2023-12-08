package templates

import (
	"bytes"
	"embed"
	"io/fs"
	"path/filepath"
	"text/template"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

//go:embed templates/config.yaml
var defaultConfig []byte

//go:embed templates
var DefaultTemplates embed.FS

type Config struct {
	Templates map[string]Template `yaml:"templates"`
}

type Template struct {
	Destination string `yaml:"dest"`
}

var DefaultConfig Config

func init() {
	err := yaml.Unmarshal(defaultConfig, &DefaultConfig)
	if err != nil {
		panic(err)
	}

	// TODO: validate defaults with schema
}

type Templates struct {
	files       fs.FS
	useDefaults bool
}

func New(files fs.FS) *Templates {
	return &Templates{
		files:       files,
		useDefaults: true,
	}
}

func (t *Templates) Render(values map[string]any) (map[string][]byte, error) {
	rendered := make(map[string][]byte)

	tmpl := template.New("").Funcs(funcMap())

	var err error
	if t.useDefaults {
		tmpl, err = tmpl.ParseFS(DefaultTemplates, "templates/*.tpl")
		if err != nil {
			return rendered, err
		}
	}

	var config Config

	if t.files != nil {
		tmpl, err = tmpl.ParseFS(t.files, "*.tpl")
		if err != nil {
			return rendered, err
		}

		content, err := fs.ReadFile(t.files, "config.yaml")
		if !t.useDefaults && err != nil {
			return rendered, err
		}

		if err == nil {
			err = yaml.Unmarshal(content, &config)
			if err != nil {
				return rendered, err
			}
		}
	}

	if t.useDefaults {
		err = mergo.Merge(&config, DefaultConfig)
		if err != nil {
			return rendered, err
		}
	}

	data := struct {
		Values map[string]any
	}{values}

	for source, templateSpec := range config.Templates {
		var buf bytes.Buffer
		err = tmpl.ExecuteTemplate(&buf, filepath.Base(source), data)
		if err != nil {
			return rendered, err
		}

		rendered[templateSpec.Destination] = buf.Bytes()
	}

	return rendered, nil
}
