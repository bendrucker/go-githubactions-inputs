package inputs

import (
	"os"
	"strings"
)

var defaultDecoder = &decoder{
	lookup: lookupInputEnv,
}

type decoder struct {
	lookup lookupInputFunc
}

// HasInput checks whether an input matching the specified name was passed.
func (d *decoder) HasInput(name string) bool {
	_, ok := d.lookup(name)
	return ok
}

// GetInput gets the input by the given name.
func (d *decoder) GetInput(name string) string {
	v, _ := d.lookup(name)
	return strings.TrimSpace(v)
}

func lookupInputEnv(name string) (string, bool) {
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ToUpper(name)
	name = "INPUT_" + name

	return os.LookupEnv(name)
}

type lookupInputFunc func(string) (string, bool)
