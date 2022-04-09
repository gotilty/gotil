package internal

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

func FilterBy(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var slice reflect.Value
		assign := false
		for i := 0; i < val.Len(); i++ {
			value := val.Index(i)
			result := f(value.Interface(), i)
			if !assign {
				t2 := value.Type()
				slice = reflect.Zero(reflect.SliceOf(t2))
				assign = true
			}
			if result {
				slice = reflect.Append(slice, value)
			}
		}

		return slice.Interface(), nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}
