package repository

import (
	"os"
	"path/filepath"
	"time"
)

type FileRepository interface {
	Save(content string) (string, error)
}

type fileRepository struct {
	storagePath string
}

func NewFileRepository(storagePath string) FileRepository {
	os.MkdirAll(storagePath, 0755)
	return &fileRepository{storagePath: storagePath}
}

func (r *fileRepository) Save(content string) (string, error) {
	filename := filepath.Join(r.storagePath, time.Now().Format("20060102-150405")+".txt")
	return filename, os.WriteFile(filename, []byte(content), 0644)
}