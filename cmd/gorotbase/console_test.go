package main

import "testing"

func Test_runConsole(t *testing.T) {
	if err := runConsole(cmdConsole, nil); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}
