package main

import "testing"

func TestCommand_Name(t *testing.T) {
	if expected, actual := "console", cmdConsole.Name(); actual != expected {
		t.Errorf("cmdConsole.Name() should be %q [actual: %q]", expected, actual)
	}
}

func TestCommand_Usage(t *testing.T) {
	cmdConsole.Usage()
}
