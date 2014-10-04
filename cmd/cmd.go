package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Cmd is an implementation of a gorot command like 'gorot create'.
type Cmd struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(*Cmd, []string)
	// UsageLine is the one-line usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string
	// Short is the short description shown in the 'gorot help' output.
	Short string
	// Long is the long message shown in the 'gorot help <this-command>' output.
	Long string
	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet
}

// Name returns the command's name: the first word in the usage line.
func (c *Cmd) Name() string {
	return strings.Split(c.UsageLine, " ")[0]
}

// Usage prints the command's usage to standard error.
func (c *Cmd) Usage() {
	fmt.Fprintf(os.Stderr, "usage: %s\n\n", c.UsageLine)
	fmt.Fprintf(os.Stderr, "%s\n", strings.TrimSpace(c.Long))
}
