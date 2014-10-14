package main

import (
	"os"
	"testing"
)

const dirnameTest = "testdir"

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
	defer os.RemoveAll(dirnameTest)

	os.RemoveAll(dirnameTest)

	if err := os.Mkdir(dirnameTest, os.ModePerm); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}

	if err := runCreate(cmdCreate, []string{dirnameTest}); err == nil {
		t.Error("error should be occurred")
	}
}

func Test_runCreate(t *testing.T) {
	defer os.RemoveAll(dirnameTest)

	os.RemoveAll(dirnameTest)

	if err := runCreate(cmdCreate, []string{dirnameTest}); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}
