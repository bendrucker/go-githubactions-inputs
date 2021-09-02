package inputs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		name   string
		inputs map[string]string
		value  interface{}
		check  func(interface{})
	}{
		{
			name:   "string",
			inputs: map[string]string{"foo": "bar"},
			value:  &struct{ Foo string }{},
			check: func(v interface{}) {
				assert.Equal(t, "bar", v.(*struct{ Foo string }).Foo)
			},
		},
		{
			name:   "int",
			inputs: map[string]string{"foo": "1"},
			value:  &struct{ Foo int }{},
			check: func(v interface{}) {
				assert.Equal(t, 1, v.(*struct{ Foo int }).Foo)
			},
		},
		{
			name:   "int64",
			inputs: map[string]string{"foo": "1"},
			value:  &struct{ Foo int64 }{},
			check: func(v interface{}) {
				assert.Equal(t, int64(1), v.(*struct{ Foo int64 }).Foo)
			},
		},
		{
			name:   "float",
			inputs: map[string]string{"foo": "1.1"},
			value:  &struct{ Foo float64 }{},
			check: func(v interface{}) {
				assert.Equal(t, 1.1, v.(*struct{ Foo float64 }).Foo)
			},
		},
		{
			name:   "bool",
			inputs: map[string]string{"foo": "true"},
			value:  &struct{ Foo bool }{},
			check: func(v interface{}) {
				assert.True(t, v.(*struct{ Foo bool }).Foo)
			},
		},
		{
			name:   "string slice",
			inputs: map[string]string{"foo": "bar1, bar2"},
			value:  &struct{ Foo []string }{},
			check: func(v interface{}) {
				assert.Equal(t, []string{"bar1", "bar2"}, v.(*struct{ Foo []string }).Foo)
			},
		},
		{
			name:   "int slice",
			inputs: map[string]string{"foo": "1, 2"},
			value:  &struct{ Foo []int }{},
			check: func(v interface{}) {
				assert.Equal(t, []int{1, 2}, v.(*struct{ Foo []int }).Foo)
			},
		},
		{
			name:   "int64 slice",
			inputs: map[string]string{"foo": "1, 2"},
			value:  &struct{ Foo []int64 }{},
			check: func(v interface{}) {
				assert.Equal(t, []int64{1, 2}, v.(*struct{ Foo []int64 }).Foo)
			},
		},
		{
			name:   "float64 slice",
			inputs: map[string]string{"foo": "1.1, 2.2"},
			value:  &struct{ Foo []float64 }{},
			check: func(v interface{}) {
				assert.Equal(t, []float64{1.1, 2.2}, v.(*struct{ Foo []float64 }).Foo)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for k, v := range tc.inputs {
				reset := setTestInput(k, v)
				t.Cleanup(func() { reset() })
			}

			err := Decode(tc.value)
			if err != nil {
				t.Fatal(err)
			}

			tc.check(tc.value)
		})
	}
}
