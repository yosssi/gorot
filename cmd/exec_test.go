package cmd

import (
	"errors"
	"flag"
	"os"
	"testing"
)

var errTest = errors.New("test error")

func init() {
	setVars()
}

func setVars() {
	a := ""

	cs := []*Cmd{
		&Cmd{
			Run:       func(cmd *Cmd, args []string) error { return nil },
			UsageLine: "version",
			Short:     "",
			Long:      "",
		},
		&Cmd{
			Run:       func(cmd *Cmd, args []string) error { return errTest },
			UsageLine: "error",
			Short:     "",
			Long:      "",
		},
	}

	ExecInit(a, cs)
}

func TestExecInit(t *testing.T) {
	a := "testApp"

	cs := []*Cmd{
		&Cmd{},
	}

	ExecInit(a, cs)

	defer setVars()

	if app != a {
		t.Errorf("app should be %q [actual: %q]", a, app)
	}

	if len(cmds) != len(cs) {
		t.Errorf("cmds should be %+v [actual: %+v]", cs, cmds)
	}

	for i, c := range cmds {
		if c != cs[i] {
			t.Errorf("cmds should be %+v [actual: %+v]", cs, cmds)
		}
	}
}

func TestExec_noArgs(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0]}

	if code := Exec(); code != codeErr {
		t.Errorf("exit code should be %d [actual: %d]", codeErr, code)
	}
}

func TestExec_helpErr(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "help", "not_exist_command"}

	if code := Exec(); code != codeErr {
		t.Errorf("exit code should be %d [actual: %d]", codeErr, code)
	}
}

func TestExec_help(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "help"}

	if code := Exec(); code != codeSuc {
		t.Errorf("exit code should be %d [actual: %d]", codeSuc, code)
	}
}

func TestExec_unknownCmd(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "not_exist_command"}

	if code := Exec(); code != codeErr {
		t.Errorf("exit code should be %d [actual: %d]", codeErr, code)
	}
}

func TestExec_cmdParseErr(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "version", "-a"}

	if code := Exec(); code != codeErr {
		t.Errorf("exit code should be %d [actual: %d]", codeErr, code)
	}
}

func TestExec_cmdRunErr(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "error"}

	if code := Exec(); code != codeErr {
		t.Errorf("exit code should be %d [actual: %d]", codeErr, code)
	}
}

func TestExec(t *testing.T) {
	resetForTesting(nil)

	os.Args = []string{os.Args[0], "version"}

	if code := Exec(); code != codeSuc {
		t.Errorf("exit code should be %d [actual: %d]", codeSuc, code)
	}
}

func Test_usage(t *testing.T) {
	usage()
}

func Test_help_argsNil(t *testing.T) {
	if err := help(nil); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}

func Test_help_manyArgs(t *testing.T) {
	if err := help([]string{"cmd1", "cmd2"}); err != errTooManyArgs {
		t.Errorf("error should be %q [actual: %q]", errTooManyArgs, err)
	}
}

func Test_help_unknownCmd(t *testing.T) {
	if err, expected := help([]string{"not_exist_command"}), errors.New("unknown help topic \"not_exist_command\""); err.Error() != expected.Error() {
		t.Errorf("error should be %q [actual: %q]", expected, err)
	}
}

func Test_help_args(t *testing.T) {
	if err := help([]string{"version"}); err != nil {
		t.Errorf("error occurred [error: %q]", err)
	}
}

func Test_writeErr(t *testing.T) {
	writeErr(errTest)
}

// resetForTesting clears all flag state and sets the usage function as directed.
// After calling ResetForTesting, parse errors in flag handling will not
// exit the program.
func resetForTesting(usage func()) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.Usage = usage
}
