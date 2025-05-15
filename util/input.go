package util

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func ReadInput(filePath string) (string, error) {
	if filePath != "" {
		b, err := os.ReadFile(filePath)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	// stdin
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}
	return input, scanner.Err()
}

func ReadDirectory(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, ".txt") {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
