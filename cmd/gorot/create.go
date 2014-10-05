package main

import (
	"fmt"

	"github.com/yosssi/gorot/cmd"
)

var cmdCreate = &cmd.Cmd{
	Run:       runCreate,
	UsageLine: "create",
	Short:     "create a deployable Gorot",
	Long:      "Create creates a deployable Gorot.",
}

// runCreate creates the create command.
func runCreate(cmd *cmd.Cmd, args []string) error {
	fmt.Println("create!")
	return nil
}
