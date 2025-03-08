package main

import (
	"fmt"
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

	_, commandPath, err := FindFileInfo(commandName)
	if err == nil {
		fmt.Printf("%s is %s\n", commandName, commandPath)
		return nil
	}

	fmt.Printf("%s: not found\n", args[0])
	return nil
}
