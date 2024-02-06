package cmd

import (
	"os"
	"path/filepath"

	"github.com/sboon-gg/svctl/internal/server"
	"github.com/spf13/cobra"
)

type serverOpts struct {
	path string
}

func newServerOpts() *serverOpts {
	return &serverOpts{
		path: ".",
	}
}

func (opts *serverOpts) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&opts.path, "path", "p", opts.path, "Path to server directory (default is current directory)")
}

func (opts *serverOpts) Path() (string, error) {
	if filepath.IsAbs(opts.path) {
		return opts.path, nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(wd, opts.path), nil
}

func (opts *serverOpts) Server() (*server.Server, error) {
	path, err := opts.Path()
	if err != nil {
		return nil, err
	}

	return server.Open(path)
}
