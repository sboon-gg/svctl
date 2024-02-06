package server

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

type Cache struct {
	PID int `yaml:"pid"`
}

func NewCache() *Cache {
	return &Cache{
		PID: -1,
	}
}

func (s *Server) Cache() (*Cache, error) {
	var cache Cache

	if _, err := os.Stat(s.dotPath(CacheFile)); os.IsNotExist(err) {
		return NewCache(), nil
	}

	content, err := os.ReadFile(s.dotPath(CacheFile))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(content, &cache)
	if err != nil {
		return nil, err
	}

	return &cache, nil
}

func (s *Server) WriteCache(cache *Cache) error {
	content, err := yaml.Marshal(cache)
	if err != nil {
		return err
	}

	return os.WriteFile(s.dotPath(CacheFile), content, 0644)
}

func (s *Server) storePID(pid int) error {
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

func (s *Server) dotPath(parts ...string) string {
	return filepath.Join(append([]string{s.Path, SvctlDir}, parts...)...)
}
