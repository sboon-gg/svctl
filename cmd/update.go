package cmd

import (
	"github.com/spf13/cobra"
)

type updateOpts struct {
	*serverOpts
	dryRun bool
	watch  bool
	values []string
}

func newUpdateOpts() *updateOpts {
	return &updateOpts{
		serverOpts: newServerOpts(),
	}
}

func init() {
	rootCmd.AddCommand(updateCmd())
}

func updateCmd() *cobra.Command {
	opts := newUpdateOpts()

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd)
		},
	}

	opts.serverOpts.AddFlags(cmd)

	return cmd
}

func (opts *updateOpts) Run(cmd *cobra.Command) error {
	return nil
}
