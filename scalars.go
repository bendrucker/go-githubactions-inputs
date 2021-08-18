package inputs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

// String returns the input value as a string. It is the same as calling githubactions.GetInput and is provided for consistency/convenience.
func String(name string) string {
	return githubactions.GetInput(name)
}

// Bool returns true if the input value is true, case-insensitive
func Bool(name string) bool {
	return strings.EqualFold(String(name), "true")
}

// Int returns the input value as an integer
func Int(name string) (int, error) {
	i, err := strconv.Atoi(String(name))

	if err != nil {
		return 0, fmt.Errorf("failed to decode input %q as int: %w", name, err)
	}

	return i, nil
}

// Int64 returns the input value as an int64
func Int64(name string) (int64, error) {
	i, err := strconv.ParseInt(String(name), 10, 64)

	if err != nil {
		return 0, fmt.Errorf("failed to decode input %q as int64: %w", name, err)
	}

	return i, nil
}
