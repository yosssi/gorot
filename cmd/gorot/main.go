package main

import (
	"os"

	"github.com/yosssi/gorot/cmd"
)

const app = "gorot"

var exit = os.Exit

// Available commands
var cmds = []*cmd.Cmd{
	cmdVersion,
}

func main() {
	cmd.ExecInit(app, cmds)
	exit(cmd.Exec())
}
