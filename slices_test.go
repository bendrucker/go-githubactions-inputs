package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringSlice(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty",
			input:    "",
			expected: []string{},
		},
		{
			name:     "one",
			input:    "one",
			expected: []string{"one"},
		},
		{
			name:     "two",
			input:    "one,two",
			expected: []string{"one", "two"},
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			reset := setTestInput("teststringslice", tc.input)
			defer reset()

			assert.Equal(t, tc.expected, StringSlice("teststringslice"))
		})
	}
}

func TestIntSlice(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []int
		err      string
	}{
		{
			name:     "empty",
			input:    "",
			expected: []int{},
		},
		{
			name:     "one",
			input:    "1",
			expected: []int{1},
		},
		{
			name:     "two",
			input:    "1,2",
			expected: []int{1, 2},
		},
		{
			name:  "invalid",
			input: "1,a",
			err:   `failed to decode entry 1 in input "testintslice" as int: strconv.Atoi: parsing "a": invalid syntax`,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			reset := setTestInput("testintslice", tc.input)
			defer reset()

			i, err := IntSlice("testintslice")

			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, i)
			}
		})
	}
}

func TestInt64Slice(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []int64
		err      string
	}{
		{
			name:     "empty",
			input:    "",
			expected: []int64{},
		},
		{
			name:     "one",
			input:    "1",
			expected: []int64{1},
		},
		{
			name:     "two",
			input:    "1,2",
			expected: []int64{1, 2},
		},
		{
			name:  "invalid",
			input: "1,a",
			err:   `failed to decode entry 1 in input "testint64slice" as int64: strconv.ParseInt: parsing "a": invalid syntax`,
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			reset := setTestInput("testint64slice", tc.input)
			defer reset()

			i, err := Int64Slice("testint64slice")

			if tc.err != "" {
				assert.EqualError(t, err, tc.err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expected, i)
			}
		})
	}
}
