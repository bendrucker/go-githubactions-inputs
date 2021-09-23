package inputs

import "reflect"

func (d *decoder) reflectValue(name string, wantType reflect.Type) (reflect.Value, error) {
	switch wantType.Kind() {
	case reflect.String:
		return reflect.ValueOf(d.String(name)), nil
	case reflect.Int:
		v, err := d.Int(name)
		return reflect.ValueOf(v), err
	case reflect.Int64:
		v, err := d.Int64(name)
		return reflect.ValueOf(v), err
	case reflect.Float64:
		v, err := d.Float64(name)
		return reflect.ValueOf(v), err
	case reflect.Bool:
		return reflect.ValueOf(d.Bool(name)), nil
	case reflect.Slice:
		return d.reflectSlice(name, wantType.Elem())
	case reflect.Ptr:
		return d.reflectPtr(name, wantType.Elem())
	default:
		panic("unsupported type")
	}
}

func (d *decoder) reflectSlice(name string, elemType reflect.Type) (reflect.Value, error) {
	switch elemType.Kind() {
	case reflect.String:
		return reflect.ValueOf(d.StringSlice(name)), nil
	case reflect.Int:
		v, err := d.IntSlice(name)
		return reflect.ValueOf(v), err
	case reflect.Int64:
		v, err := d.Int64Slice(name)
		return reflect.ValueOf(v), err
	case reflect.Float64:
		v, err := d.Float64Slice(name)
		return reflect.ValueOf(v), err
	default:
		panic("unsupported slice type")
	}
}

func (d *decoder) reflectPtr(name string, elemType reflect.Type) (reflect.Value, error) {
	switch elemType.Kind() {
	case reflect.String:
		return reflect.ValueOf(d.StringPtr(name)), nil
	case reflect.Int:
		v, err := d.IntPtr(name)
		return reflect.ValueOf(v), err
	case reflect.Int64:
		v, err := d.Int64Ptr(name)
		return reflect.ValueOf(v), err
	case reflect.Float64:
		v, err := d.Float64Ptr(name)
		return reflect.ValueOf(v), err
	case reflect.Bool:
		return reflect.ValueOf(d.BoolPtr(name)), nil
	default:
		panic("unsupported ptr type")
	}
}
