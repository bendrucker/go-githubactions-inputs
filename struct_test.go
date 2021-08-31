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
			name: "string",
			inputs: map[string]string{
				"foo": "bar",
				"baz": "qux",
			},
			value: &struct {
				Foo string
				Baz string
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo string
					Baz string
				})

				assert.Equal(t, "bar", s.Foo)
				assert.Equal(t, "qux", s.Baz)
			},
		},
		{
			name: "int",
			inputs: map[string]string{
				"foo": "1",
				"baz": "2",
			},
			value: &struct {
				Foo int
				Baz int
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo int
					Baz int
				})

				assert.Equal(t, 1, s.Foo)
				assert.Equal(t, 2, s.Baz)
			},
		},
		{
			name: "int64",
			inputs: map[string]string{
				"foo": "1",
				"baz": "2",
			},
			value: &struct {
				Foo int64
				Baz int64
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo int64
					Baz int64
				})

				assert.Equal(t, int64(1), s.Foo)
				assert.Equal(t, int64(2), s.Baz)
			},
		},
		{
			name: "float",
			inputs: map[string]string{
				"foo": "1.1",
				"baz": "2.2",
			},
			value: &struct {
				Foo float64
				Baz float64
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo float64
					Baz float64
				})

				assert.Equal(t, 1.1, s.Foo)
				assert.Equal(t, 2.2, s.Baz)
			},
		},
		{
			name: "bool",
			inputs: map[string]string{
				"foo": "true",
				"baz": "false",
			},
			value: &struct {
				Foo bool
				Baz bool
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo bool
					Baz bool
				})

				assert.True(t, s.Foo)
				assert.False(t, s.Baz)
			},
		},
		// {
		// 	name: "struct with pointer",
		// 	inputs: map[string]string{
		// 		"foo": "bar",
		// 		"baz": "qux",
		// 	},
		// 	value: &struct {
		// 		Foo *string
		// 		Baz *string
		// 	}{},
		// 	check: func(v interface{}) {
		// 		s := v.(*struct {
		// 			Foo *string
		// 			Baz *string
		// 		})

		// 		assert.Equal(t, "bar", *s.Foo)
		// 		assert.Equal(t, "qux", *s.Baz)
		// 	},
		// },
		{
			name: "string slice",
			inputs: map[string]string{
				"foo": "bar1, bar2",
				"baz": "qux1, qux2",
			},
			value: &struct {
				Foo []string
				Baz []string
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo []string
					Baz []string
				})

				assert.Equal(t, []string{"bar1", "bar2"}, s.Foo)
				assert.Equal(t, []string{"qux1", "qux2"}, s.Baz)
			},
		},
		{
			name: "int slice",
			inputs: map[string]string{
				"foo": "1, 2",
				"baz": "3, 4",
			},
			value: &struct {
				Foo []int
				Baz []int
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo []int
					Baz []int
				})

				assert.Equal(t, []int{1, 2}, s.Foo)
				assert.Equal(t, []int{3, 4}, s.Baz)
			},
		},
		{
			name: "int64 slice",
			inputs: map[string]string{
				"foo": "1, 2",
				"baz": "3, 4",
			},
			value: &struct {
				Foo []int64
				Baz []int64
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo []int64
					Baz []int64
				})

				assert.Equal(t, []int64{1, 2}, s.Foo)
				assert.Equal(t, []int64{3, 4}, s.Baz)
			},
		},
		{
			name: "float64 slice",
			inputs: map[string]string{
				"foo": "1.1, 2.2",
				"baz": "3.3, 4.4",
			},
			value: &struct {
				Foo []float64
				Baz []float64
			}{},
			check: func(v interface{}) {
				s := v.(*struct {
					Foo []float64
					Baz []float64
				})

				assert.Equal(t, []float64{1.1, 2.2}, s.Foo)
				assert.Equal(t, []float64{3.3, 4.4}, s.Baz)
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
