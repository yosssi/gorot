package main

var cmdConsole = &Command{
	Run:       runConsole,
	UsageLine: "console",
	Short:     "start the Gorot console",
	Long: `
Console starts the Gorot console.
				`,
}

var consoleI bool

func init() {
	cmdConsole.Flag.BoolVar(&consoleI, "i", false, "")
}

// runConsole runs the console command.
func runConsole(cmd *Command, args []string) {

}
