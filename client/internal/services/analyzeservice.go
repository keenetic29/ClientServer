package services

import (
    "os"
    "path/filepath"
    "client/internal/domain"
    "client/internal/transport"
)

type FileAnalyzer struct {
    client   *transport.HTTPClient
    filePath string
}

func New(client *transport.HTTPClient) *FileAnalyzer {
    return &FileAnalyzer{client: client}
}

func (a *FileAnalyzer) SetPath(path string) error {
    if _, err := os.Stat(path); err != nil {
        return domain.ErrFileNotFound
    }
    a.filePath = path
    return nil
}

func (a *FileAnalyzer) Analyze() (*domain.FileAnalysisResponse, error) {
    if a.filePath == "" {
        return nil, domain.ErrInvalidPath
    }

    content, err := os.ReadFile(a.filePath)
    if err != nil {
        return nil, err
    }

    req := domain.FileUploadRequest{
        FileName: filepath.Base(a.filePath),
        Content:  string(content),
    }

    return a.client.SendAnalysis(req)
}