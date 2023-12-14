package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"dario.cat/mergo"
	"github.com/sboon-gg/svctl/pkg/config"
	"github.com/sboon-gg/svctl/pkg/templates"
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
	si, err := newServerInstance(opts.installationPath)
	if err != nil {
		return err
	}

	if !si.HasSvctlDir() {
		return errors.New("Script has not been initialized, run `init` first.")
	}

	conf, err := si.Config()
	if err != nil {
		return err
	}

	t, err := si.Templates()

	valuesSources := conf.Values

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	for _, file := range opts.values {
		if !filepath.IsAbs(file) {
			file = filepath.Join(wd, file)
		}
		valuesSources = append(valuesSources, config.Values{
			File: &file,
		})
	}

	var allValues map[string]any

	for _, values := range valuesSources {
		if values.File != nil {
			path := *values.File
			if !filepath.IsAbs(path) {
				path = filepath.Join(si.path, path)
			}

			err = mergeValuesFile(t, &allValues, path)
			if err != nil {
				return err
			}
		} else if values.Data != nil {
			err = mergo.Merge(&allValues, values.Data, mergo.WithOverride)
			if err != nil {
				return err
			}
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

func mergeValuesFile(t *templates.Templates, allValues *map[string]any, file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	out, err := t.RenderFromString(string(content))
	if err != nil {
		return err
	}

	extraValues := make(map[string]any)
	err = yaml.Unmarshal([]byte(out), &extraValues)
	if err != nil {
		return err
	}

	return mergo.Merge(allValues, extraValues, mergo.WithOverride)
}
