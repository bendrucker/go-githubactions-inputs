package inputs

import "reflect"

func reflectValue(name string, wantType reflect.Type) (reflect.Value, error) {
	switch wantType.Kind() {
	case reflect.String:
		return reflect.ValueOf(String(name)), nil
	case reflect.Int:
		v, err := Int(name)
		return reflect.ValueOf(v), err
	case reflect.Int64:
		v, err := Int64(name)
		return reflect.ValueOf(v), err
	case reflect.Float64:
		v, err := Float64(name)
		return reflect.ValueOf(v), err
	case reflect.Bool:
		return reflect.ValueOf(Bool(name)), nil
	default:
		panic("unsupported type")
	}
}
