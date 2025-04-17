package service

import (
	"bufio"
	"strings"

	"server/internal/domain"
	"server/internal/repository"
)

type AnalyzerService interface {
	AnalyzeFile(fileName, content string) (*domain.FileAnalysis, error)
}

type analyzerService struct {
	repo repository.FileRepository
}

func NewAnalyzerService(repo repository.FileRepository) AnalyzerService {
	return &analyzerService{repo: repo}
}

func (s *analyzerService) AnalyzeFile(fileName, content string) (*domain.FileAnalysis, error) {
	lines, words, chars := 0, 0, 0
	scanner := bufio.NewScanner(strings.NewReader(content))
	
	for scanner.Scan() {
		lines++
		line := scanner.Text()
		chars += len(line)
		words += len(strings.Fields(line))
	}

	if _, err := s.repo.Save(content); err != nil {
		return nil, err
	}

	return &domain.FileAnalysis{
		FileName: fileName,
		Lines:    lines,
		Words:    words,
		Chars:    chars,
	}, nil
}