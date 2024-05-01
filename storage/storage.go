package storage

import (
	"os"
	"path/filepath"
)

type Storage struct {
	RootDir string
}

func NewStorage(rootDir string) *Storage {
	return &Storage{RootDir: rootDir}
}

func (s *Storage) SaveArtifact(path string, data []byte) error {
	fullPath := filepath.Join(s.RootDir, path)
	if err := os.MkdirAll(filepath.Dir(fullPath), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(fullPath, data, os.ModePerm)
}

func (s *Storage) GetArtifact(path string) ([]byte, error) {
	fullPath := filepath.Join(s.RootDir, path)
	return os.ReadFile(fullPath)
}
