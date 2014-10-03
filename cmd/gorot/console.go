package main

import "github.com/yosssi/gorot/cmd"

var cmdConsole = &cmd.Cmd{
	Run:       runConsole,
	UsageLine: "console",
	Short:     "start the Gorot console",
	Long: `
Console starts the Gorot console.
				`,
}

// runConsole runs the console command.
func runConsole(cmd *cmd.Cmd, args []string) {}
