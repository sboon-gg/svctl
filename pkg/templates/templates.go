package templates

import (
	"bytes"
	"io/fs"
	"os"
	"text/template"

	"dario.cat/mergo"
	"gopkg.in/yaml.v3"
)

const (
	DefaultsFileName = "defaults.yaml"
	configFileName   = "config.yaml"
)

type Values map[string]any

type Data struct {
	Values Values
}

type Template struct {
	Source      string `yaml:"src"`
	Destination string `yaml:"dest"`
	Reloadable  bool   `yaml:"reloadable"`
}

type TemplatesConfig struct {
	Templates []Template `yaml:"templates"`
}

type Renderer struct {
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

func New(config *TemplatesConfig, files fs.FS) (*Renderer, error) {
	t := &Renderer{
		config: config,
		files:  files,
	}

	err := t.loadDefaults()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewFromPath(path string) (*Renderer, error) {
	dir := os.DirFS(path)
	return NewFromFS(dir)
}

func NewFromFS(files fs.FS) (*Renderer, error) {
	conf, err := ReadConfig(files)
	if err != nil {
		return nil, err
	}

	return New(conf, files)
}

func (t *Renderer) template(name, tplContent string) (*template.Template, error) {
	return template.New(name).Funcs(FuncMap()).Parse(tplContent)
}

func (t *Renderer) prepData(values Values) (Data, error) {
	data := Data{
		Values: Values{},
	}

	err := mergo.Map(&data.Values, t.defaults)
	if err != nil {
		return data, err
	}

	err = mergo.Map(&data.Values, values, mergo.WithOverride)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (t *Renderer) Render(name, template string, values Values) ([]byte, error) {
	tmpl, err := t.template(name, template)
	if err != nil {
		return nil, err
	}

	data, err := t.prepData(values)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (t *Renderer) RenderAll(values map[string]any) (map[string][]byte, error) {
	rendered := make(map[string][]byte)

	for _, tmplSpec := range t.config.Templates {
		content, err := fs.ReadFile(t.files, tmplSpec.Source)
		if err != nil {
			return rendered, err
		}

		out, err := t.Render(tmplSpec.Source, string(content), values)
		if err != nil {
			return rendered, err
		}

		rendered[tmplSpec.Destination] = out
	}

	// Second run to allow for values interpolation in values files
	for path, content := range rendered {
		out, err := t.Render(path, string(content), values)
		if err != nil {
			return rendered, err
		}

		rendered[path] = out
	}

	return rendered, nil
}

func (t *Renderer) loadDefaults() error {
	var err error
	t.defaultsContent, err = fs.ReadFile(t.files, DefaultsFileName)
	if err != nil {
		return err
	}

	defaultsMap := make(map[string]any)

	err = yaml.Unmarshal(t.defaultsContent, &defaultsMap)
	if err != nil {
		return err
	}

	t.defaults = defaultsMap

	return nil
}

func (t *Renderer) DefaultsContent() []byte {
	return t.defaultsContent
}
