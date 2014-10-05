package main

import (
	"fmt"
	"os"
	"testing"
)

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

func Test_runCreate_mkdirErr(t *testing.T) {
	dirname := "testdir"

	defer os.Remove(dirname)

	os.Remove(dirname)

	if err := os.Mkdir(dirname, os.ModePerm); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}

	errExpected := fmt.Errorf("mkdir %s: file exists", dirname)

	if err := runCreate(cmdCreate, []string{dirname}); err == nil || err.Error() != errExpected.Error() {
		t.Errorf("err should be %q [actual: %q]", errExpected, err)
	}
}

func Test_runCreate(t *testing.T) {
	dirname := "testdir"

	defer os.Remove(dirname)

	os.Remove(dirname)

	if err := runCreate(cmdCreate, []string{dirname}); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}
