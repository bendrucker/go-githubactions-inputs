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

func TestStringPtr(t *testing.T) {
	assert.Nil(t, StringPtr("unset"))

	reset := setTestInput("teststring", "hello world")
	defer reset()

	assert.Equal(t, "hello world", *StringPtr("teststring"))
}

func TestBool(t *testing.T) {
	cases := []struct {
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

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			reset := setTestInput("testbool", tc.input)
			defer reset()

			assert.Equal(t, tc.expected, Bool("testbool"))
		})
	}
}

func TestBoolPtr(t *testing.T) {
	assert.Nil(t, BoolPtr("unset"))

	reset := setTestInput("testbool", "true")
	defer reset()

	assert.True(t, *BoolPtr("testbool"))
}

func TestInt(t *testing.T) {
	cases := []struct {
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

	for _, tc := range cases {
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

func TestIntPtr(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		nilPtr   bool
		expected int
		err      string
	}{
		{
			name:   "unset",
			input:  "",
			nilPtr: true,
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
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			d := &decoder{
				lookupenv: func(_ string) (string, bool) {
					return tc.input, tc.input != ""
				},
				getenv: func(_ string) string {
					return tc.input
				},
			}

			actual, err := d.IntPtr("testint")

			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				require.NoError(t, err)

				if tc.nilPtr {
					assert.Nil(t, actual)
				} else {
					assert.Equal(t, tc.expected, *actual)
				}
			}
		})
	}
}

func TestInt64(t *testing.T) {
	cases := []struct {
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

	for _, tc := range cases {
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

func TestFloat64(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected float64
		err      string
	}{
		{
			name:  "unset",
			input: "",
			err:   `failed to decode input "testfloat" as float64: strconv.ParseFloat: parsing "": invalid syntax`,
		},
		{
			name:     "int",
			input:    "123",
			expected: 123,
		},
		{
			name:     "float",
			input:    "12.5",
			expected: 12.5,
		},
		{
			name:  "alpha",
			input: "abc",
			err:   `failed to decode input "testfloat" as float64: strconv.ParseFloat: parsing "abc": invalid syntax`,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			reset := setTestInput("testfloat", tc.input)
			defer reset()

			actual, err := Float64("testfloat")

			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, actual)
			}
		})
	}
}
