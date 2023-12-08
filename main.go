package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"path/filepath"

	"embed"
	"text/template"

	"github.com/sboon-gg/prbf2-templates/pkg/config"
	"gopkg.in/yaml.v3"
)

// print templates
// write templates
// handle extras
// compare local templates with existing

//go:embed defaults.yaml
var defaults []byte

//go:embed templates/*
var templates embed.FS

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	content, err := templates.ReadFile("templates/config.yaml")
	if err != nil {
		return err
	}

	var conf config.Config
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return err
	}

	values := make(map[string]any)
	err = yaml.Unmarshal(defaults, &values)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Funcs(
		template.FuncMap{
			"pyBool": pyBool,
			"quote":  quote,
		},
	).ParseFS(templates, "templates/*/*.tpl")
	if err != nil {
		return err
	}

	data := struct {
		Values map[string]any
	}{values}

	s := ""
	buf := bytes.NewBufferString(s)

	for _, templateSpec := range conf.Templates {
		println("rendering", templateSpec.Source, "to", templateSpec.Destination)
		err = tmpl.ExecuteTemplate(buf, filepath.Base(templateSpec.Source), data)
		if err != nil {
			return err
		}
	}

	println(buf.String())

	return nil
}

func pyBool(v any) (string, error) {
	b, ok := v.(bool)
	if !ok {
		return "", fmt.Errorf("Expect type bool - got %T", v)
	}

	if b {
		return "True", nil
	}

	return "False", nil
}

func quote(v any) string {
	return fmt.Sprintf("%q", v)
}
