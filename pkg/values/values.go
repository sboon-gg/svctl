package values

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed defaults.yaml
var defaultValues []byte

var DefaultValues map[string]any

func init() {
	err := yaml.Unmarshal(defaultValues, &DefaultValues)
	if err != nil {
		panic(err)
	}
}
