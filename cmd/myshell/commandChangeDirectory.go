package main

import (
	"errors"
	"fmt"
	"os"
)

func commandChangeDirectory(args []string) error {
	if len(args) != 1 {
		return errors.New("Must provide a directory path")
	}
	dirPath := args[0]
	err := os.Chdir(dirPath)
	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", dirPath)
	}
	return nil
}
