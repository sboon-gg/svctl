package config

import "github.com/sboon-gg/prbf2-templates/pkg/templates"

type Config struct {
	Templates []templates.Template `yaml:"templates"`
}

var defaultConfig = Config{
	Templates: templates.DefaultTemplates,
}
