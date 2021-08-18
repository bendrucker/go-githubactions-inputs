package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	assert.Equal(t, "", String("unset"))

	reset := setTestInput("teststring", "hello world")
	defer reset()

	assert.Equal(t, "hello world", String("teststring"))
}

func TestBool(t *testing.T) {
	tc := []struct {
		name     string
		input    string
		expected bool
		err      string
	}{
		{
			name:     "unset",
			input:    "",
			expected: false,
		},
		{
			name:     "true",
			input:    "true",
			expected: true,
		},
		{
			name:     "false",
			input:    "false",
			expected: false,
		},
		{
			name:     "title",
			input:    "True",
			expected: true,
		},
		{
			name:     "caps",
			input:    "TRUE",
			expected: true,
		},
	}

	for _, tc := range tc {
		t.Run(tc.name, func(t *testing.T) {
			reset := setTestInput("testbool", tc.input)
			defer reset()

			assert.Equal(t, tc.expected, Bool("testbool"))
		})
	}
}

func TestInt(t *testing.T) {
	tc := []struct {
		name     string
		input    string
		expected int
		err      string
	}{
		{
			name:  "unset",
			input: "",
			err:   `failed to decode input "testint" as int: strconv.Atoi: parsing "": invalid syntax`,
		},
		{
			name:     "valid",
			input:    "123",
			expected: 123,
		},
		{
			name:  "float",
			input: "12.5",
			err:   `failed to decode input "testint" as int: strconv.Atoi: parsing "12.5": invalid syntax`,
		},
		{
			name:  "alpha",
			input: "abc",
			err:   `failed to decode input "testint" as int: strconv.Atoi: parsing "abc": invalid syntax`,
		},
	}

	for _, tc := range tc {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			reset := setTestInput("testint", tc.input)
			defer reset()

			actual, err := Int("testint")

			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	tc := []struct {
		name     string
		input    string
		expected int64
		err      string
	}{
		{
			name:  "unset",
			input: "",
			err:   `failed to decode input "testint" as int64: strconv.ParseInt: parsing "": invalid syntax`,
		},
		{
			name:     "valid",
			input:    "123",
			expected: 123,
		},
		{
			name:  "float",
			input: "12.5",
			err:   `failed to decode input "testint" as int64: strconv.ParseInt: parsing "12.5": invalid syntax`,
		},
		{
			name:  "alpha",
			input: "abc",
			err:   `failed to decode input "testint" as int64: strconv.ParseInt: parsing "abc": invalid syntax`,
		},
	}

	for _, tc := range tc {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			reset := setTestInput("testint", tc.input)
			defer reset()

			actual, err := Int64("testint")

			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
