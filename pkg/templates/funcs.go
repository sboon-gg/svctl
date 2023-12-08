package templates

import (
	"fmt"
	"text/template"
)

func funcMap() template.FuncMap {
	return template.FuncMap{
		"pyBool": pyBool,
		"quote":  quote,
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
