package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplatesRender(t *testing.T) {
	tests := []struct {
		name     string
		values   Values
		expected []RenderOutput
	}{
		{
			name:   "defaults",
			values: Values{},
			expected: []RenderOutput{
				{
					Template: Template{
						Source:      "file.yaml.tpl",
						Destination: "file.yaml",
					},
					Content: []byte("pyBool: True\nanother: test-string\nquoted: \"should-be-quoted\"\nnegativeBool: false\n"),
				},
			},
		},
		{
			name: "overwrite",
			values: Values{
				"boolTest": false,
				"test":     "changed-string",
				"quoted":   "but different",
			},
			expected: []RenderOutput{
				{
					Template: Template{
						Source:      "file.yaml.tpl",
						Destination: "file.yaml",
					},
					Content: []byte("pyBool: False\nanother: changed-string\nquoted: \"but different\"\nnegativeBool: false\n"),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmpl, err := NewFromPath("./testdata/example")
			assert.NoError(t, err)

			out, err := tmpl.Render(test.values)
			assert.NoError(t, err)

			assert.Len(t, out, len(test.expected))

			for i, v := range test.expected {
				assert.Equal(t, v, out[i])
			}
		})
	}
}
