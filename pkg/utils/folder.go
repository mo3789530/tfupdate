package utils

import (
	"log"
	"os"
)

type FolderUtils interface {
	ListDir() []string
}

type folder struct {
	dir string
}

func NewFolderUtils(dir string) FolderUtils {
	return &folder{
		dir: dir,
	}
}

func (f *folder) ListDir() []string {
	files, err := os.ReadDir(f.dir)
	if err != nil {
		log.Fatalf("error list dir: %s", err)
	}
	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, file.Name())
			continue
		}
	}
	return paths
}
