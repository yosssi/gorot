package cmd

import "testing"

var cmdTest = &Cmd{
	Run:       func(c *Cmd, args []string) {},
	UsageLine: "test",
	Short:     "test command",
	Long:      "This is a Gorot test command.",
}

func TestCommand_Name(t *testing.T) {
	if expected, actual := "test", cmdTest.Name(); actual != expected {
		t.Errorf("cmdConsole.Name() should be %q [actual: %q]", expected, actual)
	}
}

func TestCommand_Usage(t *testing.T) {
	cmdTest.Usage()
}
