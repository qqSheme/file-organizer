package main

import (
	"fmt"
	"os"
)

var DefaultRules = map[string]string{
	".jpg":  "Images",
	".png":  "Images",
	".jpeg": "Images",
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".txt":  "Documents",
	".mp3":  "Music",
	".wav":  "Music",
	".mp4":  "Video",
	".avi":  "Video",
	".zip":  "Archives",
	".rar":  "Archives",
}

type FileOrganizer struct {
	sourceDir      string
	rulesMap       map[string]string
	processedFiles int
	logFile        *os.File
}

func NewFileOrganizer(sourceDir string) (*FileOrganizer, error) {
	if sourceDir == "" {
		return nil, fmt.Errorf("Невалидный путь, пустая строка")
	}
	info, err := os.Stat(sourceDir)
	if err != nil {
		return nil, fmt.Errorf("Путь недоступен: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("Не явялеться путём к папке")
	}
	return &FileOrganizer{sourceDir: sourceDir}, nil
}

func main() {
	for i, v := range DefaultRules {
		fmt.Printf("%s -> %v\n", i, v)
	}
}
