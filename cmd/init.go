/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

type initOpts struct {
	installPath   string
	templatesRepo string
	token         string
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

	cmd.Flags().StringVar(&opts.templatesRepo, "templates", "", "Repository with templates")
	cmd.Flags().StringVar(&opts.token, "token", "", "Token to use when cloning templates repo")

	return cmd
}

func (opts *initOpts) Run() error {
	si, err := newServerInstance(opts.installPath)
	if err != nil {
		return err
	}

	if si.HasSvctlDir() {
		return errors.New("svctl was already initialized on this path - remove .svctl directory before initializing")
	}

	err = si.CreateDir()
	if err != nil {
		return err
	}

	if opts.templatesRepo != "" {
		err = si.CloneTemplates(opts.templatesRepo, opts.token)
		if err != nil {
			return err
		}
	}

	err = si.WriteDefaultConfig()
	if err != nil {
		return err
	}

	return si.WriteValues()
}

func init() {
	rootCmd.AddCommand(initCmd())
}
