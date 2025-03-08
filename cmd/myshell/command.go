package main

type Command struct {
	name     string
	callback func(args []string) error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"exit": {
			name:     "exit [<exit_code>]",
			callback: commandExit,
		},
		"pwd": {
			name:     "pwd",
			callback: commandPrintWorkingDirectory,
		},
		"cd": {
			name:     "cd <directory_path>",
			callback: commandChangeDirectory,
		},
		"type": {
			name:     "type <command_name>",
			callback: commandType,
		},
		"echo": {
			name:     "echo [param1, param2, ...]",
			callback: commandEcho,
		},
	}
}
