package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func commandType(args []string) error {
	if len(args) == 0 {
		return nil
	}
	commandName := args[0]
	_, isShellBuiltin := getCommands()[commandName]
	if isShellBuiltin {
		fmt.Printf("%s is a shell builtin\n", commandName)
		return nil
	}

	envPath := os.Getenv("PATH")
	pathDirs := strings.Split(envPath, ":")
	for _, dir := range pathDirs {
		commandPath := filepath.Join(dir, commandName)
		_, err := os.Stat(commandPath)
		if err == nil {
			fmt.Printf("%s is %s\n", commandName, commandPath)
			return nil
		}
	}

	fmt.Printf("%s: not found\n", args[0])
	return nil
}
