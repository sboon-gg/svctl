package prbf2update

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/hashicorp/go-version"
)

const (
	latestVersionUrl = "http://prbf2.cdn.ancientdev.com/version.json"
	patchMetaUrlFmt  = "http://prbf2.cdn.ancientdev.com/patch_%s.json"
)

var caches = make(map[string]*Cache)

type Cache struct {
	path  string
	mutex *sync.Mutex
}

func NewCache(path string) *Cache {
	path = filepath.Clean(path)

	if cache, ok := caches[path]; ok {
		return cache
	}

	cache := &Cache{
		path:  path,
		mutex: &sync.Mutex{},
	}

	caches[path] = cache

	return cache
}

func (c *Cache) FetchFor(currentVersion, requiredVersion string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	patchFilePaths, err := c.cachePatches(currentVersion, requiredVersion)
	if err != nil {
		return err
	}

	// copy patch to temp
	for _, patchFilePath := range patchFilePaths {
		patchFile := filepath.Base(patchFilePath)

		dst := os.TempDir()
		dstFilePath := filepath.Join(dst, patchFile)

		_, err := os.Stat(dstFilePath)
		if err == nil {
			continue
		}

		data, err := os.ReadFile(patchFilePath)
		if err != nil {
			return err
		}

		if err := os.WriteFile(dstFilePath, data, 0644); err != nil {
			return err
		}
	}

	return nil
}

func (c *Cache) LatestVersion() (string, error) {
	content, err := httpGet(latestVersionUrl)
	if err != nil {
		return "", err
	}

	latestVersion := struct {
		Latest string `json:"Latest"`
	}{}
	if err := json.Unmarshal(content, &latestVersion); err != nil {
		return "", fmt.Errorf("failed to decode version.json: %w", err)
	}

	return latestVersion.Latest, nil
}

func (c *Cache) ChangedFiles(fromVer, toVer string) ([]string, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	patchFilePaths, err := c.cachePatches(fromVer, toVer)
	if err != nil {
		return nil, err
	}

	var changedFiles []string

	for _, patchFilePath := range patchFilePaths {
		files, err := patchChangedFiles(patchFilePath)
		if err != nil {
			return nil, err
		}

		changedFiles = append(changedFiles, files...)
	}

	return changedFiles, nil
}

func (c *Cache) cachePatches(fromVer, toVer string) ([]string, error) {
	if toVer == fromVer {
		return nil, nil
	}

	cv, err := version.NewVersion(fromVer)
	if err != nil {
		return nil, err
	}

	var patchFilePaths []string

	for {
		println("Fetching patch info", toVer)

		info, err := c.patchMeta(toVer)
		if err != nil {
			return nil, err
		}

		println("Fetching patch from", info.Requires, "to", toVer)

		patchUrl := info.ServerData[rand.Intn(len(info.ServerData))]
		patchPath, err := c.cachePatch(patchUrl)
		if err != nil {
			return nil, err
		}

		patchFilePaths = append(patchFilePaths, patchPath)

		toVer = info.Requires

		rv, err := version.NewVersion(toVer)
		if err != nil {
			return nil, err
		}

		if rv.LessThan(cv) {
			return nil, fmt.Errorf("No patch available for version %s", fromVer)
		}

		if rv.Equal(cv) {
			break
		}
	}

	return patchFilePaths, nil
}

func (c *Cache) cachePatch(patchUrl string) (string, error) {
	filename, err := filenameFromURL(patchUrl)
	if err != nil {
		return "", err
	}

	patchPath := filepath.Join(c.path, filename)

	if _, err := os.Stat(patchPath); err == nil {
		// patch already exists
		return patchPath, nil
	}

	resp, err := http.Get(patchUrl)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch %s: %s", patchUrl, resp.Status)
	}

	defer resp.Body.Close()

	file, err := os.Create(patchPath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		return "", err
	}

	return patchPath, nil
}

type patchMeta struct {
	Requires   string   `json:"Requires"`
	ServerData []string `json:"ServerData"`
	// There are more fields in the actual JSON
	// but we don't need them
}

func (c *Cache) patchMeta(version string) (*patchMeta, error) {
	path := filepath.Join(c.path, fmt.Sprintf("patch_%s.json", version))
	data, err := os.ReadFile(path)
	if err != nil {
		url := fmt.Sprintf(patchMetaUrlFmt, version)
		data, err = httpGet(url)
		if err != nil {
			return nil, err
		}

		_ = os.WriteFile(path, data, 0644)
	}

	var info patchMeta
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, fmt.Errorf("failed to decode patch_%s.json: %w", version, err)
	}

	return &info, nil
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch %s: %s", url, resp.Status)
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func filenameFromURL(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	parts := strings.Split(u.Path, "/")
	return parts[len(parts)-1], nil
}
