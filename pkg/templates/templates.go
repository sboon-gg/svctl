package templates

import (
	"bytes"
	"embed"
	"io/fs"
	"path/filepath"
	"text/template"
)

//go:embed templates
var templateFiles embed.FS

type Template struct {
	Source      string `yaml:"src"`
	Destination string `yaml:"dest"`
}

var DefaultTemplates = []Template{
	{
		Source:      "serversettings.con.tpl",
		Destination: "mods/pr/settings/serversettings.con",
	},
	{
		Source:      "realityconfig_admin.py.tpl",
		Destination: "mods/pr/python/game/realityconfig_admin.py",
	},
	{
		Source:      "realityconfig_common.py.tpl",
		Destination: "mods/pr/python/game/realityconfig_common.py",
	},
	{
		Source:      "realityconfig_private.py.tpl",
		Destination: "mods/pr/python/game/realityconfig_private.py",
	},
	{
		Source:      "realityconfig_tracker.py.tpl",
		Destination: "mods/pr/python/game/realityconfig_tracker.py",
	},
}

type Templates struct {
	files  fs.FS
	config []Template
}

func New(config []Template, files fs.FS) *Templates {
	return &Templates{
		config: config,
		files:  files,
	}
}

func (t *Templates) Render(values map[string]any) (map[string][]byte, error) {
	rendered := make(map[string][]byte)

	tmpl := template.New("").Funcs(funcMap())

	tmpl, err := tmpl.ParseFS(templateFiles, "templates/*.tpl")
	if err != nil {
		return rendered, err
	}

	if t.files != nil {
		tmpl, err = tmpl.ParseFS(t.files, "*.tpl")
		if err != nil {
			return rendered, err
		}
	}

	data := make(map[string]any)

	data["Values"] = values

	for _, templateSpec := range t.config {
		var buf bytes.Buffer
		err = tmpl.ExecuteTemplate(&buf, filepath.Base(templateSpec.Source), data)
		if err != nil {
			return rendered, err
		}

		rendered[templateSpec.Destination] = buf.Bytes()
	}

	return rendered, nil
}
