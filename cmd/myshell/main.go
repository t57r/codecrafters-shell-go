package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func IsShellBuiltin(s string) bool {
	builtins := []string{"echo", "exit", "type"}
	return slices.Contains(builtins, s)
}

func main() {
	for true {
		// get input
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("Error while reading input %v\n", err)
			os.Exit(1)
		}
		cmd = strings.TrimSuffix(cmd, "\n")

		tokens := strings.Fields(cmd)
		if len(tokens) == 2 && tokens[0] == "exit" {
			if exitCode, err := strconv.Atoi(tokens[1]); err == nil {
				os.Exit(exitCode)
			}
		}

		if len(tokens) > 2 && tokens[0] == "echo" {
			fmt.Println(strings.Join(tokens[1:], " "))
			continue
		}

		if len(tokens) == 2 && tokens[0] == "type" {
			if IsShellBuiltin(tokens[1]) {
				fmt.Printf("%s is a shell builtin\n", tokens[1])
			} else {
				fmt.Printf("%s: not found\n", tokens[1])
			}
			continue
		}

		// at this point command is not verified
		fmt.Printf("%s: command not found\n", cmd)
	}
}
