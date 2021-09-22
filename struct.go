package inputs

import (
	"reflect"
)

// Decode  reads the inputs and stores them in the struct pointed to by v
func Decode(v interface{}) error {
	if v == nil {
		panic("Decode: v is nil")
	}

	val := reflect.ValueOf(v)

	if val.Kind() != reflect.Ptr {
		panic("Decode: v is not a pointer")
	}

	for i := 0; i < val.Elem().NumField(); i++ {
		fieldVal := val.Elem().Field(i)
		field := val.Elem().Type().Field(i)

		val, err := reflectValue(field.Name, fieldVal.Type())
		if err != nil {
			return err
		}

		fieldVal.Set(val)
	}

	return nil
}
