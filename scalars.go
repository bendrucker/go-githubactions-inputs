package inputs

import (
	"fmt"
	"strconv"
	"strings"
)

// String returns the input value as a string. It is the same as calling githubactions.GetInput and is provided for consistency/convenience.
func (d *decoder) String(name string) string {
	return d.GetInput(name)
}

// String returns the input value as a string. It is the same as calling githubactions.GetInput and is provided for consistency/convenience.
func String(name string) string {
	return defaultDecoder.String(name)
}

// StringPtr returns the input value as a string pointer. If no matching input is found, it returns nil.
func (d *decoder) StringPtr(name string) *string {
	if !d.HasInput(name) {
		return nil
	}

	v := d.String(name)
	return &v
}

// StringPtr returns the input value as a string pointer. If no matching input is found, it returns nil.
func StringPtr(name string) *string {
	return defaultDecoder.StringPtr(name)
}

// Bool returns true if the input value is true, case-insensitive
func (d *decoder) Bool(name string) bool {
	return strings.EqualFold(d.String(name), "true")
}

// Bool returns true if the input value is true, case-insensitive
func Bool(name string) bool {
	return defaultDecoder.Bool(name)
}

// BoolPtr returns the input value as a bool pointer. If no matching input is found, it returns nil.
func (d *decoder) BoolPtr(name string) *bool {
	if !d.HasInput(name) {
		return nil
	}

	v := d.Bool(name)
	return &v
}

// BoolPtr returns the input value as a string pointer. If no matching input is found, it returns nil.
func BoolPtr(name string) *bool {
	return defaultDecoder.BoolPtr(name)
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

// IntPtr returns the input value as an int pointer. If no matching input is found, it returns nil. If the input value is not an integer, it returns an error.
func (d *decoder) IntPtr(name string) (*int, error) {
	if !d.HasInput(name) {
		return nil, nil
	}

	i, err := d.Int(name)
	return &i, err
}

// IntPtr returns the input value as an int pointer. If no matching input is found, it returns nil. If the input value is not an integer, it returns an error.
func IntPtr(name string) (*int, error) {
	return defaultDecoder.IntPtr(name)
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

// Int64Ptr returns the input value as an int pointer. If no matching input is found, it returns nil. If the input value is not an integer, it returns an error.
func (d *decoder) Int64Ptr(name string) (*int64, error) {
	if !d.HasInput(name) {
		return nil, nil
	}

	i, err := d.Int64(name)
	return &i, err
}

// Int64Ptr returns the input value as an int pointer. If no matching input is found, it returns nil. If the input value is not an integer, it returns an error.
func Int64Ptr(name string) (*int64, error) {
	return defaultDecoder.Int64Ptr(name)
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

// Float64Ptr returns the input value as an int pointer. If no matching input is found, it returns nil. If the input value is not an integer, it returns an error.
func (d *decoder) Float64Ptr(name string) (*float64, error) {
	if !d.HasInput(name) {
		return nil, nil
	}

	i, err := d.Float64(name)
	return &i, err
}

// Float64Ptr returns the input value as an int pointer. If no matching input is found, it returns nil. If the input value is not an integer, it returns an error.
func Float64Ptr(name string) (*float64, error) {
	return defaultDecoder.Float64Ptr(name)
}
