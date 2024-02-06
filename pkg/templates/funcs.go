package templates

import (
	"fmt"
	"os"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func FuncMap() template.FuncMap {
	f := sprig.TxtFuncMap()

	extra := template.FuncMap{
		"pyBool": pyBool,
		"quote":  quote,
		"env":    env,
	}

	for k, v := range extra {
		f[k] = v
	}

	return f
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
