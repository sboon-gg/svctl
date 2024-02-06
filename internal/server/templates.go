package server

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/goccy/go-yaml"
	"github.com/sboon-gg/svctl/pkg/templates"
)

func (s *Server) Templates() (*templates.Renderer, error) {
	return templates.NewFromPath(s.dotPath(TemplatesDir))
}

func (s *Server) Values() (templates.Values, error) {
	content, err := os.ReadFile(filepath.Join(s.dotPath(ValuesFile)))
	if err != nil {
		return nil, err
	}

	var values templates.Values
	err = yaml.Unmarshal(content, &values)
	if err != nil {
		return nil, err
	}

	return values, nil
}

func cloneTemplates(path, repoURL, token string) error {
	var auth transport.AuthMethod
	if token != "" {
		auth = &http.BasicAuth{
			Username: "git",
			Password: token,
		}
	}

	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:             repoURL,
		Auth:            auth,
		InsecureSkipTLS: true,
	})
	return err
}

func writeValues(path string, content []byte) error {
	commented := commentOutWholeYamlFile(string(content))

	return os.WriteFile(filepath.Join(path, ValuesFile), []byte(commented), 0755)
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
