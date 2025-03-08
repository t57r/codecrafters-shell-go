package main

import (
	"os"
	"strconv"
)

func commandExit(args []string) error {
	exitCode := 0
	if len(args) > 0 {
		exitCode, _ = strconv.Atoi(args[0])
	}
	os.Exit(exitCode)
	return nil
}
