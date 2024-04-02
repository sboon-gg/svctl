package prbf2update

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	updaterPath = "mods/pr/bin"
	modDescPath = "mods/pr/mod.desc"
)

type Result struct {
	ChangedFiles []string
	OldVersion   string
	NewVersion   string
}

type PRBF2Update struct {
	path  string
	cache *Cache
}

func New(path string, cache *Cache) *PRBF2Update {
	if cache == nil {
		tempDir, err := os.MkdirTemp("", "svctl-update-cache-*")
		if err != nil {
			tempDir = os.TempDir()
		}
		cache = NewCache(tempDir)
	}

	return &PRBF2Update{
		path:  path,
		cache: cache,
	}
}

func (u *PRBF2Update) Update() (*Result, error) {
	old, err := u.currentVersion()
	if err != nil {
		return nil, err
	}

	latest, err := u.cache.LatestVersion()
	if err != nil {
		return nil, err
	}

	err = u.cache.FetchFor(old, latest)
	if err != nil {
		return nil, err
	}

	out, err := u.update()
	if err != nil {
		fmt.Println(string(out))
		return nil, err
	}

	new, err := u.currentVersion()
	if err != nil {
		return nil, err
	}

	changedFiles, err := u.cache.ChangedFiles(old, new)
	if err != nil {
		return nil, err
	}

	return &Result{
		ChangedFiles: changedFiles,
		OldVersion:   old,
		NewVersion:   new,
	}, nil
}

func (u *PRBF2Update) currentVersion() (string, error) {
	content, err := os.ReadFile(filepath.Join(u.path, modDescPath))
	if err != nil {
		return "", err
	}

	var modDesc struct {
		XMLName xml.Name `xml:"mod"`
		Version string   `xml:"version"`
	}
	err = xml.Unmarshal(content, &modDesc)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(modDesc.Version), nil
}
