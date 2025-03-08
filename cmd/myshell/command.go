package main

type Command struct {
	name     string
	callback func(args []string) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			name:     "exit",
			callback: commandExit,
		},
		"pwd": {
			name:     "pwd",
			callback: commandPrintWorkingDirectory,
		},
		"type": {
			name:     "type",
			callback: commandType,
		},
		"echo": {
			name:     "echo",
			callback: commandEcho,
		},
	}
}
