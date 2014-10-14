package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/yosssi/gorot/cmd"
)

var cmdCreate = &cmd.Cmd{
	Run:       runCreate,
	UsageLine: "create",
	Short:     "create a deployable Gorot",
	Long:      "Create creates a deployable Gorot.",
}

// Errors
var (
	errCreateNameNotSpecified = errors.New("gorot name is not specified")
	errCreateTooManyArgs      = errors.New("too many arguments given")
)

// runCreate creates the create command.
func runCreate(cmd *cmd.Cmd, args []string) error {
	l := len(args)

	switch {
	case l < 1:
		return errCreateNameNotSpecified
	case l > 1:
		return errCreateTooManyArgs
	}

	dirname := args[0]

	if err := os.Mkdir(dirname, os.ModePerm); err != nil {
		return err
	}

	tmplData := map[string]interface{}{
		"App": dirname,
	}

	for _, filename := range AssetNames() {
		// Asset does not return an error because filename comes from AssetNames.
		data, _ := Asset(filename)

		tmpl, err := template.New(filename).Delims("[[", "]]").Parse(string(data))

		if err != nil {
			return err
		}

		bf := new(bytes.Buffer)

		if err := tmpl.Execute(bf, tmplData); err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(dirname, filename), bf.Bytes(), os.ModePerm); err != nil {
			return err
		}
	}

	fmt.Println("create!")
	return nil
}
