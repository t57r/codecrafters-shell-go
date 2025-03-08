package main

import (
	"fmt"
	"os"
)

func commandPrintWorkingDirectory(args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error printing working dir: %s\n", err)
	}
	fmt.Println(pwd)
	return nil
}
