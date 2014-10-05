package main

import (
	"fmt"

	"github.com/yosssi/gorot/cmd"
)

var cmdConsole = &cmd.Cmd{
	Run:       runConsole,
	UsageLine: "console",
	Short:     "launch a Gorot in the console mode",
	Long:      "Launch a Gorot in the console mode.",
}

// runConsole creates the console command.
func runConsole(cmd *cmd.Cmd, args []string) error {
	fmt.Println("console!!")
	return nil
}
