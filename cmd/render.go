package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"dario.cat/mergo"
	"github.com/sboon-gg/svctl/pkg/config"
	"github.com/sboon-gg/svctl/pkg/templates"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type renderOpts struct {
	installationPath string
	defaults         bool
	dryRun           bool
	watch            bool
	values           []string
}

func init() {
	rootCmd.AddCommand(renderCmd())
}

func renderCmd() *cobra.Command {
	var opts renderOpts

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
	cmd.Flags().BoolVar(&opts.watch, "watch", false, "Watch all values and config files")
	cmd.Flags().StringSliceVar(&opts.values, "values", []string{}, "")

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

	si, err := newServerInstance(opts.installationPath)
	if err != nil {
		return files, err
	}

	if !si.HasSvctlDir() {
		return files, errors.New("Script has not been initialized, run `init` first.")
	}

	files = append(files, si.path)
	files = append(files, filepath.Join(si.path, templatesDir))

	conf, err := si.Config()
	if err != nil {
		return files, err
	}

	t, err := si.Templates()

	valuesSources := conf.Values

	wd, err := os.Getwd()
	if err != nil {
		return files, err
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
				return files, err
			}
		} else if values.Data != nil {
			err = mergo.Merge(&allValues, values.Data, mergo.WithOverride)
			if err != nil {
				return files, err
			}
		}
	}

	out, err := t.RenderAll(allValues)
	if err != nil {
		return files, err
	}

	if opts.dryRun {
		for name, content := range out {
			fmt.Printf(`File: %s\n---\n%s`, name, string(content))
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
		path := filepath.Join(opts.installationPath, name)

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

func mergeValuesFile(t *templates.Templates, allValues *map[string]any, file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	out, err := t.Render(file, string(content), *allValues)
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
