package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/yosssi/gorot/cmd"
)

var cmdCreate = &cmd.Cmd{
	Run:       runCreate,
	UsageLine: "create",
	Short:     "create a deployable Gorot",
	Long:      "Create creates a deployable Gorot.",
}

// Errors
var (
	errCreateNameNotSpecified = errors.New("gorot name is not specified")
	errCreateTooManyArgs      = errors.New("too many arguments given")
)

// runCreate creates the create command.
func runCreate(cmd *cmd.Cmd, args []string) error {
	l := len(args)

	switch {
	case l < 1:
		return errCreateNameNotSpecified
	case l > 1:
		return errCreateTooManyArgs
	}

	name := args[0]

	if err := os.Mkdir(name, os.ModePerm); err != nil {
		return err
	}

	fmt.Println("create!")
	return nil
}
