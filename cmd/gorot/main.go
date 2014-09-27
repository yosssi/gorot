package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

var exit = os.Exit

// exit codes
const (
	exitCodeError       = 1
	exitCodeNotExecuted = 2
)

// template strings
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

// templates
var (
	usageTpl = template.Must(template.New("usage").Parse(usageTplStr))
	helpTpl  = template.Must(template.New("help").Funcs(
		template.FuncMap{
			"trim": strings.TrimSpace,
		},
	).Parse(helpTplStr))
)

// Commands lists the available commands.
// The order here is the order in which they are printed by 'gorot help'.
var commands = []*Command{
	cmdConsole,
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		if err := usageTpl.Execute(os.Stderr, commands); err != nil {
			writeErr(err)
			exit(exitCodeError)
			return
		}

		exit(exitCodeNotExecuted)
		return
	}

	if args[0] == "help" {
		if err := help(args[1:]); err != nil {
			writeErr(err)
			exit(exitCodeError)
			return
		}

		exit(exitCodeNotExecuted)
		return
	}
}

// help implements the 'help' command.
func help(args []string) error {
	switch l := len(args); {
	case l == 0:
		return usageTpl.Execute(os.Stderr, commands)
	case l != 1:
		return errors.New("too many arguments given")
	}

	cmdName := args[0]

	for _, cmd := range commands {
		if cmd.Name() == cmdName {
			return helpTpl.Execute(os.Stdout, cmd)
		}
	}

	return fmt.Errorf("unknown help topic %#q", cmdName)
}

// writeErr writes the error to standard error.
func writeErr(err error) {
	write(os.Stderr, err.Error())
}

// write writes s to w.
func write(w io.Writer, s string) {
	w.Write([]byte(s + "\n"))
}
