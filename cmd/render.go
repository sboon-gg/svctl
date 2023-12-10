package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"dario.cat/mergo"
	"github.com/sboon-gg/prbf2-templates/pkg/templates"
	"github.com/sboon-gg/prbf2-templates/pkg/values"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type renderFlagOpts struct {
	installationPath string
	defaults         bool
	dryRun           bool
	values           []string
}

func renderCmd() *cobra.Command {
	var opts renderFlagOpts

	cmd := &cobra.Command{
		Use:          "render",
		Args:         cobra.ExactArgs(1),
		Short:        "A brief description of your command",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.installationPath = args[0]

			return opts.Run(cmd)
		},
	}

	// cmd.Flags().BoolVar(&opts.defaults, "defaults", true, "Use default values")
	cmd.Flags().BoolVar(&opts.dryRun, "dry-run", false, "Print out rendered files")
	cmd.Flags().StringSliceVar(&opts.values, "values", []string{}, "")

	return cmd
}

func (opts *renderFlagOpts) Run(cmd *cobra.Command) error {
	// si, err := newServerInstance(opts.installationPath)
	// if err != nil {
	// 	return err
	// }
	//
	// if !si.HasConfigCache() {
	// 	return errors.New("Script has not been initialized, run `init` first.")
	// }

	t := templates.New(templates.DefaultTemplates, templates.DefaultTemplateFiles)

	var allValues map[string]any
	err := mergo.Map(&allValues, values.DefaultValues)
	if err != nil {
		return err
	}

	for _, valuesFile := range opts.values {
		content, err := os.ReadFile(valuesFile)
		if err != nil {
			return err
		}

		tmpl := template.New("").Funcs(templates.FuncMap())
		tmpl, err = tmpl.Parse(string(content))
		if err != nil {
			return err
		}

		var buf bytes.Buffer
		err = tmpl.Execute(&buf, allValues)
		if err != nil {
			return err
		}

		var extraValues map[string]any
		err = yaml.Unmarshal(buf.Bytes(), &extraValues)
		if err != nil {
			return err
		}

		err = mergo.Map(&allValues, extraValues)
		if err != nil {
			return err
		}
	}

	out, err := t.Render(allValues)
	if err != nil {
		return err
	}

	for name, content := range out {
		if opts.dryRun {
			fmt.Printf(`File: %s
---
%s
`, name, string(content))
		} else {
			path := filepath.Join(opts.installationPath, name)

			err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				return err
			}

			err = os.WriteFile(path, content, 0755)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(renderCmd())
}
