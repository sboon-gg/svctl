package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sboon-gg/prbf2-templates/pkg/templates"
	"github.com/sboon-gg/prbf2-templates/pkg/values"
	"github.com/spf13/cobra"
)

type renderFlagOpts struct {
	installationPath string
	defaults         bool
	dryRun           bool
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

	cmd.Flags().BoolVar(&opts.defaults, "defaults", false, "Use default values")
	cmd.Flags().BoolVar(&opts.dryRun, "dry-run", false, "Print out rendered files")

	return cmd
}

func (opts *renderFlagOpts) Run(cmd *cobra.Command) error {
	si, err := newServerInstance(opts.installationPath)
	if err != nil {
		return err
	}

	if !si.HasConfigCache() {
		return errors.New("Script has not been initialized, run `init` first.")
	}

	t := templates.New(templates.DefaultTemplates, nil)

	out, err := t.Render(values.DefaultValues)
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
