package inputs

import (
	"fmt"
	"strconv"
	"strings"
)

// StringSlice returns the input value as a slice of strings, splitting on commas.
// Each entry is stripped of leading and trailing whitespace.
func StringSlice(name string) []string {
	v := String(name)
	if v == "" {
		return []string{}
	}

	s := strings.Split(v, ",")

	for i, str := range s {
		s[i] = strings.TrimSpace(str)
	}

	return s
}

// IntSlice returns the input value as a slice of ints, splitting on commas.
// If any entry is not an integer, an error is returned.
func IntSlice(name string) ([]int, error) {
	s := StringSlice(name)
	ints := make([]int, len(s))

	for i, str := range s {
		v, err := strconv.Atoi(str)

		if err != nil {
			return nil, fmt.Errorf("failed to decode entry %d in input %q as int: %w", i, name, err)
		}

		ints[i] = v
	}

	return ints, nil
}

// Int64Slice returns the input value as a slice of int64s, splitting on commas.
// If any entry is not an integer, an error is returned.
func Int64Slice(name string) ([]int64, error) {
	s := StringSlice(name)
	ints := make([]int64, len(s))

	for i, str := range s {
		v, err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			return nil, fmt.Errorf("failed to decode entry %d in input %q as int64: %w", i, name, err)
		}

		ints[i] = v
	}

	return ints, nil
}
