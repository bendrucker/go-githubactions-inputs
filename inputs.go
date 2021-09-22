package inputs

import (
	"os"
	"strings"
)

var defaultDecoder = &decoder{
	getenv:    os.Getenv,
	lookupenv: os.LookupEnv,
}

type decoder struct {
	getenv    getenvFunc
	lookupenv lookupEnvFunc
}

// HasInput gets the input by the given name.
func (d *decoder) HasInput(name string) bool {
	_, ok := d.lookupenv(inputVar(name))
	return ok
}

// GetInput gets the input by the given name.
func (d *decoder) GetInput(name string) string {
	return strings.TrimSpace(d.getenv(inputVar(name)))
}

func inputVar(name string) string {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ToUpper(name)
	name = "INPUT_" + name

	return name
}

type getenvFunc func(string) string
type lookupEnvFunc func(string) (string, bool)
