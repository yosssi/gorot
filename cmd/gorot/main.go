package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

var exit = os.Exit

// Exit codes
const (
	exitCodeSuccess     = 0
	exitCodeError       = 1
	exitCodeNotExecuted = 2
)

// Template strings
const (
	usageTplStr = `Gorot is a pure Go customizable robot.

Usage:

	gorot command [arguments]

The commands are:
{{range .}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "gorot help [command]" for more information about a command.

`

	helpTplStr = `usage: gorot {{.UsageLine}}

{{.Long | trim}}
`
)

// Templates
var (
	usageTpl = template.Must(template.New("usage").Parse(usageTplStr))
	helpTpl  = template.Must(template.New("help").Funcs(
		template.FuncMap{
			"trim": strings.TrimSpace,
		},
	).Parse(helpTplStr))
)

// Command list
var commands = []*Command{
	cmdConsole,
}

// Errors
var (
	errTooManyArgs = errors.New("too many arguments given")
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		// Ignore the error which is never returned.
		usageTpl.Execute(os.Stderr, commands)

		exit(exitCodeNotExecuted)

		return
	}

	cmdName := args[0]

	if cmdName == "help" {
		if err := help(args[1:]); err != nil {
			writeErr(err)
			exit(exitCodeError)
			return
		}

		exit(exitCodeNotExecuted)
		return
	}

	for _, cmd := range commands {
		if cmd.Name() != cmdName {
			continue
		}

		cmd.Flag.Usage = func() { cmd.Usage() }

		cmd.Flag.Parse(args[1:])

		args = cmd.Flag.Args()

		cmd.Run(cmd, args)

		exit(exitCodeSuccess)
		return
	}

	writeErr(fmt.Errorf("gorot: unknown subcommand %q", cmdName))
	exit(exitCodeError)
}

// help implements the 'help' command.
func help(args []string) error {
	switch l := len(args); {
	case l == 0:
		return usageTpl.Execute(os.Stderr, commands)
	case l != 1:
		return errTooManyArgs
	}

	cmdName := args[0]

	for _, cmd := range commands {
		if cmd.Name() == cmdName {
			return helpTpl.Execute(os.Stdout, cmd)
		}
	}

	return fmt.Errorf("unknown help topic %q", cmdName)
}

// writeErr writes the error to standard error.
func writeErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}
