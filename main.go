package main

import (
	"os"

	"github.com/sboon-gg/prbf2-templates/pkg/config"
	"gopkg.in/yaml.v3"
)

// print templates
// write templates
// handle extras
//

func main() {
	err := run()
	if err != nil {
		panic(err)
	}
}

func run() error {
	content, err := os.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	var conf config.Config
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return err
	}

	return nil
}
