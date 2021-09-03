package inputs

import (
	"reflect"
)

// Decode  reads the inputs and stores them in the struct pointed to by v
func Decode(v interface{}) error {
	val := reflect.ValueOf(v)

	if val.IsNil() {
		panic("invalid input: nil")
	}

	if val.Kind() != reflect.Ptr {
		panic("invalid input: non-pointer")
	}

	for i := 0; i < val.Elem().NumField(); i++ {
		fieldVal := val.Elem().Field(i)
		field := val.Elem().Type().Field(i)

		switch kind := fieldVal.Kind(); kind {
		case reflect.String:
			fieldVal.SetString(String(field.Name))
		case reflect.Int, reflect.Int64:
			intV, err := Int64(field.Name)
			if err != nil {
				return &DecodeError{}
			}

			fieldVal.SetInt(intV)
		case reflect.Float64:
			floatV, err := Float64(field.Name)
			if err != nil {
				return &DecodeError{}
			}

			fieldVal.SetFloat(floatV)
		case reflect.Bool:
			fieldVal.SetBool(Bool(field.Name))
		case reflect.Slice:
			switch elementKind := fieldVal.Type().Elem().Kind(); elementKind {
			case reflect.String:
				fieldVal.Set(reflect.ValueOf(StringSlice(field.Name)))
			case reflect.Int:
				intSlice, err := IntSlice(field.Name)
				if err != nil {
					return &DecodeError{}
				}

				fieldVal.Set(reflect.ValueOf(intSlice))
			case reflect.Int64:
				intSlice, err := Int64Slice(field.Name)
				if err != nil {
					return &DecodeError{}
				}

				fieldVal.Set(reflect.ValueOf(intSlice))
			case reflect.Float64:
				floatSlice, err := Float64Slice(field.Name)
				if err != nil {
					return &DecodeError{}
				}

				fieldVal.Set(reflect.ValueOf(floatSlice))
			default:
				panic("invalid input: unknown slice elem type: " + elementKind.String())
			}
		default:
			panic("invalid input: unknown type: " + kind.String())
		}
	}

	return nil
}

type DecodeError struct{}

func (e DecodeError) Error() string {
	return "decode error"
}
