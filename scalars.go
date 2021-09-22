package inputs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

// String returns the input value as a string. It is the same as calling githubactions.GetInput and is provided for consistency/convenience.
func (d *decoder) String(name string) string {
	return githubactions.GetInput(name)
}

// String returns the input value as a string. It is the same as calling githubactions.GetInput and is provided for consistency/convenience.
func String(name string) string {
	return defaultDecoder.String(name)
}

// Bool returns true if the input value is true, case-insensitive
func (d *decoder) bool(name string) bool {
	return strings.EqualFold(d.String(name), "true")
}

// Bool returns true if the input value is true, case-insensitive
func Bool(name string) bool {
	return defaultDecoder.bool(name)
}

// Int returns the input value as an integer
func (d *decoder) Int(name string) (int, error) {
	i, err := strconv.Atoi(d.String(name))

	if err != nil {
		return 0, fmt.Errorf("failed to decode input %q as int: %w", name, err)
	}

	return i, nil
}

// Int returns the input as an integer
func Int(name string) (int, error) {
	return defaultDecoder.Int(name)
}

// Int64 returns the input value as an int64
func (d *decoder) Int64(name string) (int64, error) {
	i, err := strconv.ParseInt(d.String(name), 10, 64)

	if err != nil {
		return 0, fmt.Errorf("failed to decode input %q as int64: %w", name, err)
	}

	return i, nil
}

// Int64 returns the input as an int64
func Int64(name string) (int64, error) {
	return defaultDecoder.Int64(name)
}

// Float64 returns the input value as a float64
func (d *decoder) Float64(name string) (float64, error) {
	f, err := strconv.ParseFloat(d.String(name), 64)

	if err != nil {
		return 0, fmt.Errorf("failed to decode input %q as float64: %w", name, err)
	}

	return f, nil
}

// Float64 returns the input as a float64
func Float64(name string) (float64, error) {
	return defaultDecoder.Float64(name)
}
