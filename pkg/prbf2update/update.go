package prbf2update

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

const (
	updaterPath = "mods/pr/bin"
	modDescPath = "mods/pr/mod.desc"
)

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

func (u *PRBF2Update) Update() (string, string, error) {
	old, err := u.currentVersion()
	if err != nil {
		return "", "", err
	}

	if u.cache != nil {
		latest, err := u.cache.LatestVersion()
		if err != nil {
			return "", "", err
		}

		err = u.cache.FetchFor(old, latest)
		if err != nil {
			return "", "", err
		}
	}

	out, err := u.update()
	if err != nil {
		fmt.Println(string(out))
		return "", "", err
	}

	new, err := u.currentVersion()

	return old, new, nil
}

func (u *PRBF2Update) currentVersion() (string, error) {
	content, err := os.ReadFile(filepath.Join(u.path, modDescPath))
	if err != nil {
		return "", err
	}

	var modDesc struct {
		Mod struct {
			Version string `xml:"version"`
		} `xml:"mod"`
	}

	err = xml.Unmarshal(content, &modDesc)
	if err != nil {
		return "", err
	}

	return modDesc.Mod.Version, nil
}