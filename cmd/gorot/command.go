package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// A Command is an implementation of a gorot command
// like gorot console or gorot server.
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string)
	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string
	// Short is the short description shown in the 'gorot help' output.
	Short string
	// Long is the long message shown in the 'gorot help <this-command>' output.
	Long string
	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags bool
}

// Name returns the command's name: the first word in the usage line.
func (c *Command) Name() string {
	return strings.Split(c.UsageLine, " ")[0]
}

// Usage prints the command's usage to standard error.
func (c *Command) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
}
