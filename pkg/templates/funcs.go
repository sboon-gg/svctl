package templates

import (
	"fmt"
	"os"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"pyBool": pyBool,
		"quote":  quote,
		"env":    env,
	}
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

func env(v any) (string, error) {
	s, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("Expect type bool - got %T", v)
	}

	return os.Getenv(s), nil
}
