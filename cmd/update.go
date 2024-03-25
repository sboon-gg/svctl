package cmd

import (
	"path/filepath"

	"github.com/sboon-gg/svctl/pkg/prbf2update"
	"github.com/spf13/cobra"
)

type updateOpts struct {
	path string
}

func updateCmd() *cobra.Command {
	opts := &updateOpts{}

	cmd := &cobra.Command{
		Use: "update",
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.path = args[0]
			return opts.Run(cmd)
		},
	}

	return cmd
}

func (o *updateOpts) Run(cmd *cobra.Command) error {
	cache := prbf2update.NewCache(filepath.Join(o.path, ".update-cache"))

	u := prbf2update.New(o.path, cache)

	old, new, err := u.Update()
	if err != nil {
		return err
	}

	cmd.Println("Updated from", old, "to", new)

	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd())
}
