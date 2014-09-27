package main

import (
	"io"
	"os"
	"text/template"
)

var exit = os.Exit

// Commands lists the available commands.
// The order here is the order in which they are printed by 'gorot help'.
var commands = []*Command{
	cmdConsole,
}

var usageTplStr = `Gorot is a pure Go customizable robot.

Usage:

	gorot command [arguments]

The commands are:
{{range .}}
    {{.Name | printf "%-11s"}} {{.Short}}{{end}}

Use "gorot help [command]" for more information about a command.

`

func main() {
	if err := usage(); err != nil {
		writeErr(err)
		exit(1)
		return
	}
}

// usage prints the command usage.
func usage() error {
	t, err := template.New("usage").Parse(usageTplStr)

	if err != nil {
		return err
	}

	return t.Execute(os.Stderr, commands)
}

// writeErr writes the error to standard error.
func writeErr(err error) {
	write(os.Stderr, err.Error())
}

// write writes s to w.
func write(w io.Writer, s string) {
	w.Write([]byte(s + "\n"))
}
