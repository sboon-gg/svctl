package templates

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

const (
	DefaultsFileName = "defaults.yaml"
	configFileName   = "config.yaml"
	fileHeaderFmt    = `
###
# Default file: %s
###
`
)

type Values map[string]any

type Template struct {
	Source      string `yaml:"src"`
	Destination string `yaml:"dest"`
	Reloadable  bool   `yaml:"reloadable"`
}

type TemplatesConfig struct {
	Templates []Template `yaml:"templates"`
	Defaults  []string   `yaml:"defaults"`
}

type Templates struct {
	config          *TemplatesConfig
	files           fs.FS
	defaultsContent []byte
	defaults        Values
}

func ReadConfig(dir fs.FS) (*TemplatesConfig, error) {
	content, err := fs.ReadFile(dir, configFileName)
	if err != nil {
		return nil, err
	}

	var conf TemplatesConfig
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func New(config *TemplatesConfig, files fs.FS) (*Templates, error) {
	t := &Templates{
		config: config,
		files:  files,
	}

	err := t.loadDefaults()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewFromPath(path string) (*Templates, error) {
	dir := os.DirFS(path)

	conf, err := ReadConfig(dir)
	if err != nil {
		return nil, err
	}

	return New(conf, dir)
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

func (t *Templates) loadDefaults() error {
	defaultsContent, err := mergeDefaultsFiles(t.files, t.config.Defaults)
	if err != nil {
		return err
	}

	t.defaultsContent = defaultsContent

	defaultsMap := make(map[string]any)

	err = yaml.Unmarshal(defaultsContent, &defaultsMap)
	if err != nil {
		return err
	}

	t.defaults = defaultsMap

	return nil
}

func mergeDefaultsFiles(dir fs.FS, files []string) ([]byte, error) {
	var defaults bytes.Buffer

	for _, defaultsFile := range files {
		content, err := fs.ReadFile(dir, defaultsFile)
		if err != nil {
			return nil, err
		}

		_, err = defaults.WriteString(fmt.Sprintf(fileHeaderFmt, defaultsFile))
		if err != nil {
			return nil, err
		}

		_, err = defaults.Write(content)
		if err != nil {
			return nil, err
		}
	}

	return defaults.Bytes(), nil
}
