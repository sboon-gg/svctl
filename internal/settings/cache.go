package settings

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Cache struct {
	PID           int      `yaml:"pid"`
	UpdatePatches []string `yaml:"update_patches"`
}

func NewCache() *Cache {
	return &Cache{
		PID: -1,
	}
}

func (s *Settings) Cache() (*Cache, error) {
	var cache Cache

	cacheFile := filepath.Join(s.path, CacheFile)

	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		return NewCache(), nil
	}

	content, err := os.ReadFile(cacheFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &cache)
	if err != nil {
		return nil, err
	}

	return &cache, nil
}

func (s *Settings) WriteCache(cache *Cache) error {
	content, err := yaml.Marshal(cache)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(s.path, CacheFile), content, 0644)
}

func (s *Settings) StorePID(pid int) error {
	cache, err := s.Cache()
	if err != nil {
		return err
	}

	cache.PID = pid

	err = s.WriteCache(cache)
	if err != nil {
		return err
	}

	return nil
}
