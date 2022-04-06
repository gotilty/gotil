package collection

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

func FilterByIterate(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var slice reflect.Value
		var kind reflect.Kind
		assign := false
		for i := 0; i < val.Len(); i++ {
			value := val.Index(i)
			result := f(value.Interface(), i)
			if !assign {
				t2 := value.Type()
				kind = t2.Kind()
				slice = reflect.Zero(reflect.SliceOf(t2))
				assign = true
			}
			if result {
				slice = reflect.Append(slice, value)
			}
		}
		switch kind {
		case reflect.Int:
			return slice.Interface().([]int), nil
		case reflect.Int16:
			return slice.Interface().([]int16), nil
		case reflect.Int32:
			return slice.Interface().([]int32), nil
		case reflect.Int64:
			return slice.Interface().([]int64), nil
		case reflect.String:
			return slice.Interface().([]string), nil
		case reflect.Uint:
			return slice.Interface().([]uint), nil
		case reflect.Uint16:
			return slice.Interface().([]uint16), nil
		case reflect.Uint32:
			return slice.Interface().([]uint32), nil
		case reflect.Uint64:
			return slice.Interface().([]uint64), nil
		case reflect.Float32:
			return slice.Interface().([]float32), nil
		case reflect.Float64:
			return slice.Interface().([]float64), nil
		}
		return slice.Interface(), nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}
