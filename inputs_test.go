package inputs

import (
	"os"
	"strings"
)

// setTestInput sets the input and returns a function that will reset it to its previous value.
func setTestInput(name, value string) func() {
	// https://github.com/sethvargo/go-githubactions/blob/431bd3a55929219e916bc7aba512ec684c7210c9/actions.go#L183-L189
	key := strings.ReplaceAll(name, " ", "_")
	key = strings.ToUpper(key)
	key = "INPUT_" + key

	prev := os.Getenv(key)
	os.Setenv(key, value)

	return func() {
		os.Setenv(key, prev)
	}
}
