package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

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

		// handle command
		fmt.Printf("%s: command not found\n", cmd)
	}
}
