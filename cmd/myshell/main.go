package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("Error while reading input %v\n", err)
			os.Exit(1)
		}

		words := clearInput(cmd)
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var args = []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

		_, _, err = FindFileInfo(commandName)
		if err == nil {
			cmd := exec.Command(commandName, args...)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Print(string(stdout))
			}
			continue
		}

		fmt.Printf("%s: command not found\n", commandName)
	}
}

func clearInput(text string) []string {
	cmd := strings.TrimSuffix(text, "\n")
	return strings.Fields(cmd)
}
