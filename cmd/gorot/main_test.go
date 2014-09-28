package main

import (
	"errors"
	"testing"
)

func Test_help_lenArgsEqualsZero(t *testing.T) {
	if err := help(nil); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}

func Test_help_lenArgsNotEqualOne(t *testing.T) {
	err := help([]string{"test1", "test2"})

	if err == nil {
		t.Error("error should be returned")
	}

	if err != errTooManyArgs {
		t.Errorf("error should be %q [actual: %q]", errTooManyArgs, err)
	}
}

func Test_help_unknown(t *testing.T) {
	err := help([]string{"not_exist_cmd"})

	if err == nil {
		t.Error("error should be returned")
	}

	if actual, expected := err.Error(), "unknown help topic `not_exist_cmd`"; actual != expected {
		t.Errorf("error should be %q [actual: %q]", expected, actual)
	}
}

func Test_help(t *testing.T) {
	if err := help([]string{"console"}); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}

func Test_writeErr(t *testing.T) {
	writeErr(errors.New("test error"))
}
