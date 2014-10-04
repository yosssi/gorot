package cmd

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"text/template"
)

// Exit codes
const (
	codeSuc = 0
	codeErr = 1
)

// Template strings
const (
	strTmplUsage = `{{.app}} is a pure Go customizable robot.

Usage:

  {{.app}} command [arguments]

The commands are:
{{range .cmds}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "{{.app}} help [command]" for more information about a command.

`

	strTmplHelp = `usage: {{.app}} {{.cmd.UsageLine}}

{{.cmd.Long | trim}}
`
)

// Templates
var (
	tmplUsage = template.Must(template.New("usage").Parse(strTmplUsage))
	tmplHelp  = template.Must(template.New("help").Funcs(
		template.FuncMap{
			"trim": strings.TrimSpace,
		},
	).Parse(strTmplHelp))
)

var app string
var muApp = new(sync.RWMutex)

var cmds []*Cmd
var muCmds = new(sync.RWMutex)

// Errors
var (
	errTooManyArgs = errors.New("too many arguments given")
)

// ExecInit initializes the global variables for the execution.
func ExecInit(a string, cs []*Cmd) {
	muApp.Lock()
	defer muApp.Unlock()

	muCmds.Lock()
	defer muCmds.Unlock()

	app = a
	cmds = cs
}

// Exec parses the arguments and executes the specified command.
func Exec() int {
	// Parse the flags.
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		usage()
		return codeErr
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	if cmdName == "help" {
		if err := help(cmdArgs); err != nil {
			writeErr(err)
			return codeErr
		}

		return codeSuc
	}

	muCmds.RLock()
	defer muCmds.RUnlock()

	for _, cmd := range cmds {
		if cmd.Name() != cmdName {
			continue
		}

		cmd.Flag.Usage = cmd.Usage

		if err := cmd.Flag.Parse(cmdArgs); err != nil {
			return codeErr
		}

		if err := cmd.Run(cmd, cmd.Flag.Args()); err != nil {
			writeErr(err)
			cmd.Usage()
			return codeErr
		}

		return codeSuc
	}

	writeErr(fmt.Errorf("unknown gorot command %q", cmdName))
	return codeErr
}

// usage prints the usages to standard error.
func usage() {
	muApp.RLock()
	defer muApp.RUnlock()

	muCmds.RLock()
	defer muCmds.RUnlock()

	if tmplUsage == nil {
		return
	}

	tmplUsage.Execute(os.Stderr, map[string]interface{}{
		"app":  app,
		"cmds": cmds,
	})
}

// help implements the 'help' command.
func help(args []string) error {
	muApp.RLock()
	defer muApp.RUnlock()

	muCmds.RLock()
	defer muCmds.RUnlock()

	switch l := len(args); {
	case l == 0:
		usage()
		return nil
	case l != 1:
		return errTooManyArgs
	}

	cmdName := args[0]

	for _, cmd := range cmds {
		if cmd.Name() == cmdName {
			return tmplHelp.Execute(os.Stdout, map[string]interface{}{
				"app": app,
				"cmd": cmd,
			})
		}
	}

	return fmt.Errorf("unknown help topic %q", cmdName)
}

// writeErr writes the error to standard error.
func writeErr(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
}
