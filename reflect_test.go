package inputs

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_reflectValue(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantType  reflect.Type
		want      reflect.Value
		wantError string
	}{
		{
			name:     "string",
			wantType: reflect.TypeOf(""),
			input:    "foo",
			want:     reflect.ValueOf("foo"),
		},
		{
			name:     "int",
			wantType: reflect.TypeOf(1),
			input:    "1",
			want:     reflect.ValueOf(1),
		},
		{
			name:     "int64",
			wantType: reflect.TypeOf(int64(1)),
			input:    "1",
			want:     reflect.ValueOf(int64(1)),
		},
		{
			name:     "float",
			wantType: reflect.TypeOf(float64(1)),
			input:    "1",
			want:     reflect.ValueOf(float64(1)),
		},
		{
			name:     "bool",
			wantType: reflect.TypeOf(true),
			input:    "true",
			want:     reflect.ValueOf(true),
		},
		{
			name:      "error",
			wantType:  reflect.TypeOf(1),
			input:     "foo",
			wantError: "failed to decode",
		},
		{
			name:     "string slice",
			wantType: reflect.TypeOf([]string{}),
			input:    "foo,bar",
			want:     reflect.ValueOf([]string{"foo", "bar"}),
		},
		{
			name:     "int slice",
			wantType: reflect.TypeOf([]int{}),
			input:    "1,2",
			want:     reflect.ValueOf([]int{1, 2}),
		},
		{
			name:     "int64 slice",
			wantType: reflect.TypeOf([]int64{}),
			input:    "1,2",
			want:     reflect.ValueOf([]int64{1, 2}),
		},
		{
			name:     "float slice",
			wantType: reflect.TypeOf([]float64{}),
			input:    "1.1,2.2",
			want:     reflect.ValueOf([]float64{1.1, 2.2}),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			decoder := newTestDecoder(map[string]string{"reflected": tc.input})

			got, err := decoder.reflectValue("reflected", tc.wantType)
			if tc.wantError == "" {
				require.NoError(t, err)
			} else {
				require.Contains(t, err.Error(), tc.wantError)
				return
			}

			assert.Equal(t, tc.wantType, got.Type())
			assert.Equal(t, tc.want.Interface(), got.Interface())
		})
	}
}
