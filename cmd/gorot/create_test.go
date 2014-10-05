package main

import "testing"

func Test_runCreate_nameNotSpecified(t *testing.T) {
	if err := runCreate(cmdCreate, nil); err != errCreateNameNotSpecified {
		t.Errorf("err should be %q [actual: %q]", errCreateNameNotSpecified, err)
	}
}

func Test_runCreate_tooManyArgs(t *testing.T) {
	if err := runCreate(cmdCreate, []string{"test1", "test2"}); err != errCreateTooManyArgs {
		t.Errorf("err should be %q [actual: %q]", errCreateTooManyArgs, err)
	}
}

func Test_runCreate(t *testing.T) {
	if err := runCreate(cmdCreate, []string{"test"}); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}
