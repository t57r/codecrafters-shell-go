package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func FindFileInfo(commandName string) (os.FileInfo, string, error) {
	envPath := os.Getenv("PATH")
	pathDirs := strings.Split(envPath, ":")
	for _, dir := range pathDirs {
		commandPath := filepath.Join(dir, commandName)
		fileInfo, err := os.Stat(commandPath)
		if err == nil {
			return fileInfo, commandPath, nil
		}
	}

	return nil, "", errors.New("FileInfo has not been found")
}
