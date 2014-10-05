package main

import (
	"os"

	"github.com/yosssi/gorot/cmd"
)

// Exit function
var exit = os.Exit

// Available commands
var cmds = []*cmd.Cmd{
	cmdConsole,
}

func main() {
	// Initialize the settings for the execution.
	cmd.ExecInit(app, cmds)

	// Execute a command and exit.
	exit(cmd.Exec())
}
