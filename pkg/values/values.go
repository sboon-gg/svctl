package values

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed defaults.yaml
var DefaultValuesFile []byte

var DefaultValues map[string]any

func init() {
	err := yaml.Unmarshal(DefaultValuesFile, &DefaultValues)
	if err != nil {
		panic(err)
	}
}
