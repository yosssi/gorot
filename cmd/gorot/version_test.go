package main

import "testing"

func Test_runVersion_withArgs(t *testing.T) {
	if err := runVersion(cmdVersion, []string{"test"}); err != errVersionArgs {
		t.Errorf("err should be %q [actual: %q]", errVersionArgs, err)
	}
}

func Test_runVersion(t *testing.T) {
	if err := runVersion(cmdVersion, nil); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}
