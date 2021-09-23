package inputs

import (
	"fmt"
	"strconv"
	"strings"
)

// StringSlice returns the input value as a slice of strings, splitting on commas.
// Each entry is stripped of leading and trailing whitespace.
func (d *decoder) StringSlice(name string) []string {
	v := d.String(name)
	if v == "" {
		return []string{}
	}

	s := strings.Split(v, ",")

	for i, str := range s {
		s[i] = strings.TrimSpace(str)
	}

	return s
}

// StringSlice returns the input value as a slice of strings, splitting on commas.
// Each entry is stripped of leading and trailing whitespace.
func StringSlice(name string) []string {
	return defaultDecoder.StringSlice(name)
}

// IntSlice returns the input value as a slice of ints, splitting on commas.
// If any entry is not an integer, an error is returned.
func (d *decoder) IntSlice(name string) ([]int, error) {
	s := d.StringSlice(name)
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

// IntSlice returns the input value as a slice of strings, splitting on commas.
// Each entry is stripped of leading and trailing whitespace.
func IntSlice(name string) ([]int, error) {
	return defaultDecoder.IntSlice(name)
}

// Int64Slice returns the input value as a slice of int64s, splitting on commas.
// If any entry is not an integer, an error is returned.
func (d *decoder) Int64Slice(name string) ([]int64, error) {
	s := d.StringSlice(name)
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

// Int64Slice returns the input value as a slice of strings, splitting on commas.
// Each entry is stripped of leading and trailing whitespace.
func Int64Slice(name string) ([]int64, error) {
	return defaultDecoder.Int64Slice(name)
}

// Float64Slice returns the input value as a slice of float64s, splitting on commas.
// If any entry is not a float, an error is returned.
func (d *decoder) Float64Slice(name string) ([]float64, error) {
	s := d.StringSlice(name)
	floats := make([]float64, len(s))

	for i, str := range s {
		v, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, fmt.Errorf("failed to decode entry %d in input %q as float64: %w", i, name, err)
		}

		floats[i] = v
	}

	return floats, nil
}

// Float64Slice returns the input value as a slice of strings, splitting on commas.
// Each entry is stripped of leading and trailing whitespace.
func Float64Slice(name string) ([]float64, error) {
	return defaultDecoder.Float64Slice(name)
}
