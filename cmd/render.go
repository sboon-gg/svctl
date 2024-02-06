package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"github.com/sboon-gg/svctl/internal/server"
	"github.com/spf13/cobra"
)

type renderOpts struct {
	*serverOpts
	defaults bool
	dryRun   bool
	watch    bool
	values   []string
}

func newRenderOpts() *renderOpts {
	return &renderOpts{
		serverOpts: newServerOpts(),
	}
}

func init() {
	rootCmd.AddCommand(renderCmd())
}

func renderCmd() *cobra.Command {
	opts := newRenderOpts()

	cmd := &cobra.Command{
		Use:          "render",
		Short:        "Render templates",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd)
		},
	}

	opts.serverOpts.AddFlags(cmd)

	// cmd.Flags().BoolVar(&opts.defaults, "defaults", true, "Use default values")
	cmd.Flags().BoolVar(&opts.dryRun, "dry-run", false, "Print out rendered files")
	cmd.Flags().BoolVar(&opts.watch, "watch", false, "Watch all values and config files")
	cmd.Flags().StringSliceVar(&opts.values, "values", []string{}, "Additional values files - relative to execution working directory")

	return cmd
}

func (opts *renderOpts) Run(cmd *cobra.Command) error {
	watchedFiles, err := opts.render()
	if err != nil {
		return err
	}

	if !opts.watch {
		return nil
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	for _, file := range watchedFiles {
		err = watcher.Add(file)
		if err != nil {
			return err
		}
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}

			if event.Has(fsnotify.Write | fsnotify.Create) {
				watchedFiles, err = opts.render()
				if err != nil {
					log.Printf("Error rendering files: %s", err)
				} else {
					for _, file := range watchedFiles {
						watcher.Add(file)
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.Println("error:", err)
		}
	}
}

func (opts *renderOpts) render() ([]string, error) {
	var files []string

	path, err := opts.Path()
	if err != nil {
		return files, err
	}

	si, err := server.Open(path)
	if err != nil {
		return files, errors.New("Script has not been initialized, run `init` first.")
	}

	t, err := si.Templates()
	if err != nil {
		return files, err
	}

	allValues, err := si.Values()
	if err != nil {
		return files, err
	}

	out, err := t.RenderAll(allValues)
	if err != nil {
		return files, err
	}

	if opts.dryRun {
		for name, content := range out {
			fmt.Printf("File: %s\n---\n%s", name, string(content))
		}
	} else {
		err = opts.writeOutput(out)
		if err != nil {
			return files, err
		}
	}

	return files, nil
}

func (opts *renderOpts) writeOutput(out map[string][]byte) error {
	for name, content := range out {
		path := filepath.Join(opts.path, name)

		err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(path, content, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
