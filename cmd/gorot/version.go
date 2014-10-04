package main

import (
	"errors"
	"fmt"

	"github.com/yosssi/gorot"
	"github.com/yosssi/gorot/cmd"
)

var cmdVersion = &cmd.Cmd{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "print Gorot version",
	Long:      "Version prints the Gorot version.",
}

var errVersionArgs = errors.New(`args are specified but the "version" command does not take them`)

// runVersion runs the version command.
func runVersion(cmd *cmd.Cmd, args []string) error {
	if len(args) != 0 {
		return errVersionArgs
	}

	fmt.Printf("Gorot %s\n", gorot.Version)

	return nil
}
