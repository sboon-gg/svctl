package config

type Config struct {
	Values []Values `yaml:"values"`
}

type Values struct {
	File *string        `yaml:"file,omitempty"`
	Data map[string]any `yaml:"data,omitempty"`
}

var (
	DefaultValuesFile = "values.yaml"
	DefaultConfig     = Config{
		Values: []Values{
			{
				File: &DefaultValuesFile,
			},
		},
	}
)
