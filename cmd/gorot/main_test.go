package main

import (
	"errors"
	"flag"
	"os"
	"testing"
)

func init() {
	exit = func(_ int) {}
}

func Test_main_argsLenLessThanOne(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0]}

	main()
}

func Test_main_helpErr(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "help", "not_exist_command"}

	main()
}

func Test_main_help(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "help", "version"}

	main()
}

func Test_main_unknownSubcommand(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "not_exist_command"}

	main()
}

func Test_main(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "version"}

	main()

	cmdVersion.Flag.Usage()
}

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

	if actual, expected := err.Error(), `unknown help topic "not_exist_cmd"`; actual != expected {
		t.Errorf("error should be %q [actual: %q]", expected, actual)
	}
}

func Test_help(t *testing.T) {
	if err := help([]string{"version"}); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}

func Test_writeErr(t *testing.T) {
	writeErr(errors.New("test error"))
}

// resetForTesting clears all flag state and sets the usage function as directed.
// After calling ResetForTesting, parse errors in flag handling will not
// exit the program.
func resetForTesting(usage func()) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = usage
}
