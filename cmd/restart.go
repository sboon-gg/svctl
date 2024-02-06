package cmd

import (
	"github.com/spf13/cobra"
)

type restartOpts struct {
	*serverOpts
}

func newRestartOpts() *restartOpts {
	return &restartOpts{
		serverOpts: newServerOpts(),
	}
}

func restartCmd() *cobra.Command {
	opts := newRestartOpts()

	cmd := &cobra.Command{
		Use:          "restart",
		Short:        "Restarts the server",
		Long:         `Restarts the server`,
		SilenceUsage: true,
		RunE:         opts.Run,
	}

	opts.AddFlags(cmd)

	return cmd
}

func (o *restartOpts) AddFlags(cmd *cobra.Command) {
	o.serverOpts.AddFlags(cmd)
}

func (o *restartOpts) Run(cmd *cobra.Command, args []string) error {
	return nil
}

func init() {
	rootCmd.AddCommand(restartCmd())
}
