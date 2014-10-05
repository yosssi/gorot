package main

import "testing"

func Test_runCreate(t *testing.T) {
	if err := runCreate(cmdCreate, nil); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}
