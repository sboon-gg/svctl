package templates

import (
	"bytes"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

const (
	DefaultsFileName = ".defaults.yaml"
	configFileName   = ".config.yaml"
)

type Template struct {
	Source      string `yaml:"src"`
	Destination string `yaml:"dest"`
	Reloadable  bool   `yaml:"reloadable"`
}

type TemplatesConfig struct {
	Templates []Template `yaml:"templates"`
}

type Templates struct {
	config   TemplatesConfig
	files    fs.FS
	defaults map[string]any
}

func New(config TemplatesConfig, files fs.FS) *Templates {
	return &Templates{
		config: config,
		files:  files,
	}
}

func NewFromPath(path string) (*Templates, error) {
	t := Templates{}

	dir := os.DirFS(path)
	t.files = dir

	content, err := fs.ReadFile(dir, configFileName)
	if err != nil {
		return nil, err
	}

	var conf TemplatesConfig
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	t.config = conf

	content, err = fs.ReadFile(dir, DefaultsFileName)
	if err == nil {
		var defaults map[string]any
		err = yaml.Unmarshal(content, &defaults)
		if err != nil {
			return nil, err
		}

		t.defaults = defaults
	}

	return &t, nil
}

func (t *Templates) RenderFromString(text string) (string, error) {
	tmpl := template.New("").Funcs(FuncMap())

	data := map[string]any{
		"Values": t.defaults,
	}

	tmpl, err := tmpl.Parse(text)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (t *Templates) Render(values map[string]any) (map[string][]byte, error) {
	rendered := make(map[string][]byte)

	tmpl := template.New("").Funcs(FuncMap())

	if t.defaults != nil {
		err := mergo.Merge(&values, t.defaults)
		if err != nil {
			return rendered, err
		}
	}

	data := map[string]any{
		"Values": values,
	}

	var err error
	for _, tmplSpec := range t.config.Templates {
		tmpl, err = tmpl.ParseFS(t.files, tmplSpec.Source)
		if err != nil {
			return rendered, err
		}

		var buf bytes.Buffer
		err = tmpl.ExecuteTemplate(&buf, filepath.Base(tmplSpec.Source), data)
		if err != nil {
			return rendered, err
		}

		rendered[tmplSpec.Destination] = buf.Bytes()
	}

	// Second run to allow for values interpolation in values files
	for path, content := range rendered {
		tmpl := template.New("").Funcs(FuncMap())
		tmpl, err := tmpl.Parse(string(content))
		if err != nil {
			return rendered, err
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, data)
		if err != nil {
			return rendered, err
		}

		rendered[path] = buf.Bytes()
	}

	return rendered, nil
}
