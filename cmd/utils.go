package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/sboon-gg/svctl/pkg/config"
	"github.com/sboon-gg/svctl/pkg/templates"
	"gopkg.in/yaml.v3"
)

const (
	svctlDir   = ".svctl"
	valuesFile = "values.yaml"
	configFile = "config.yaml"

	templatesDir = "templates"
)

type serverInstance struct {
	path string
}

func newServerInstance(path string) (*serverInstance, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	return &serverInstance{
		path: filepath.Join(path, svctlDir),
	}, nil
}

func (si *serverInstance) CloneTemplates(repoURL, token string) error {
	_, err := git.PlainClone(filepath.Join(si.path, templatesDir), false, &git.CloneOptions{
		URL: repoURL,
		Auth: &http.BasicAuth{
			Username: "git",
			Password: token,
		},
		InsecureSkipTLS: true,
	})
	return err
}

func (si *serverInstance) Templates() (*templates.Templates, error) {
	return templates.NewFromPath(filepath.Join(si.path, templatesDir))
}

func (si *serverInstance) HasSvctlDir() bool {
	if _, err := os.Stat(si.path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (si *serverInstance) Config() (*config.Config, error) {
	content, err := os.ReadFile(filepath.Join(si.path, configFile))
	if err != nil {
		return nil, err
	}

	var conf config.Config
	err = yaml.Unmarshal(content, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

func (si *serverInstance) CreateDir() error {
	return os.Mkdir(si.path, os.ModePerm)
}

func (si *serverInstance) WriteDefaultConfig() error {
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	err := enc.Encode(config.DefaultConfig)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(si.path, "config.yaml"), buf.Bytes(), 0755)
}

func (si *serverInstance) WriteValues() error {
	content, err := os.ReadFile(filepath.Join(si.path, templatesDir, templates.DefaultsFileName))
	if err == nil {
		content = []byte(commentOutWholeYamlFile(string(content)))
	}

	return os.WriteFile(filepath.Join(si.path, valuesFile), content, 0755)
}

func commentOutWholeYamlFile(content string) string {
	reader := strings.NewReader(content)
	scanner := bufio.NewScanner(reader)

	builder := strings.Builder{}

	for scanner.Scan() {
		text := scanner.Text()

		if !strings.HasPrefix(strings.TrimSpace(text), "#") {
			if len(strings.TrimSpace(text)) > 0 {
				text = fmt.Sprintf("# %s", text)
			}
		}

		builder.WriteString(text)
		builder.WriteString("\n")
	}

	return builder.String()
}
