package updater

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"

	"github.com/sboon-gg/svctl/internal/server"
	"github.com/sboon-gg/svctl/pkg/prbf2"
	"github.com/sboon-gg/svctl/pkg/templates"
)

type Updater struct {
	server *server.Server
}

func New(s *server.Server) *Updater {
	return &Updater{
		server: s,
	}
}

func (u *Updater) Update() error {
	changedFiles, err := prbf2.Update(u.server.Path)
	if err != nil {
		return err
	}

	if len(changedFiles) == 0 {
		return nil
	}

	patches, err := u.prepareDiffs(changedFiles)
	if err != nil {
		return err
	}

	cache, err := u.server.Settings.Cache()
	if err != nil {
		return err
	}

	for _, patch := range patches {
		cache.UpdatePatches = append(cache.UpdatePatches, patch)
	}

	err = u.server.Settings.WriteCache(cache)
	if err != nil {
		return err
	}

	return nil
}

func (u *Updater) prepareDiffs(changedFiles []string) ([]string, error) {
	out, err := u.server.Settings.Templates.Render(templates.Values{})
	if err != nil {
		return nil, err
	}

	files := make(map[string][2][]byte)

	for _, changed := range changedFiles {
		for _, rendered := range out {
			if filepath.Clean(rendered.Destination) == filepath.Clean(changed) {
				content, err := os.ReadFile(filepath.Join(u.server.Path, changed))
				if err != nil {
					return nil, err
				}

				files[rendered.Destination] = [2][]byte{content, rendered.Content}
			}
		}
	}

	return PreparePatches(files)
}

func PreparePatches(files map[string][2][]byte) ([]string, error) {
	mem := memory.NewStorage()
	fs := memfs.New()

	repo, err := git.Init(mem, fs)
	if err != nil {
		return nil, err
	}

	for file, contents := range files {
		dirs := filepath.Dir(file)
		err := fs.MkdirAll(dirs, os.ModePerm)
		if err != nil {
			return nil, err
		}

		f, err := fs.OpenFile(file, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return nil, err
		}

		_, err = f.Write(contents[0])
		if err != nil {
			return nil, err
		}
	}

	w, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	err = w.AddGlob("*")
	if err != nil {
		return nil, err
	}

	originalHash, err := w.Commit("Original", &git.CommitOptions{})

	for file, contents := range files {
		f, err := fs.OpenFile(file, os.O_TRUNC|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return nil, err
		}

		_, err = f.Write(contents[1])
		if err != nil {
			return nil, err
		}
	}

	changedHash, err := w.Commit("Changed", &git.CommitOptions{
		All: true,
	})
	if err != nil {
		return nil, err
	}

	origCommit, err := repo.CommitObject(originalHash)
	if err != nil {
		return nil, err
	}

	changedCommit, err := repo.CommitObject(changedHash)
	if err != nil {
		return nil, err
	}

	origTree, err := origCommit.Tree()
	if err != nil {
		return nil, err
	}

	changedTree, err := changedCommit.Tree()
	if err != nil {
		return nil, err
	}

	diff, err := origTree.Diff(changedTree)
	if err != nil {
		return nil, err
	}

	patches := make([]string, 0)

	for _, d := range diff {
		_, err := d.Action()
		if err != nil {
			continue
		}

		patch, err := d.Patch()
		if err != nil {
			continue
		}

		patches = append(patches, patch.String())
	}

	return patches, nil
}
