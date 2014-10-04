package main

import "testing"

func Test_runVersion_withArgs(t *testing.T) {
	runVersion(cmdVersion, []string{"test"})
}

func Test_runVersion(t *testing.T) {
	runVersion(cmdVersion, nil)
}
