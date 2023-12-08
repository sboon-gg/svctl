/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

type initOpts struct {
	installPath string
}

func initCmd() *cobra.Command {
	var opts initOpts

	cmd := &cobra.Command{
		Use:   "init",
		Short: "A brief description of your command",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.installPath = args[0]

			return opts.Run()
		},
	}

	return cmd
}

func (opts *initOpts) Run() error {
	si, err := newServerInstance(opts.installPath)
	if err != nil {
		return err
	}

	if !si.HasConfigCache() {
		err = si.WriteDefaultConfig()
		if err != nil {
			return err
		}

		err = si.WriteDefaultValues()
		if err != nil {
			return err
		}

		return si.WriteDefaultTemplates()
	}

	return nil
}

func init() {
	rootCmd.AddCommand(initCmd())
}
