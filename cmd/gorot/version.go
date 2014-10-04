package main

import (
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

// runVersion runs the version command.
func runVersion(cmd *cmd.Cmd, args []string) {
	if len(args) != 0 {
		cmd.Usage()
		return
	}

	fmt.Printf("Gorot %s\n", gorot.Version)
}
