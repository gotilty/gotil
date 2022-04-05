package array

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

// Creates an array of values by running each element in the given array thru iteratee.
// The value to be iterated should be given as the first parameter.
// The second parameter will be a function that will take the parameter and return value type interface{}.
// Example: func square(a interface{}, i int()) interface{} {
// 	b, _ := converter.ToInt64(a)
// 	return (b * b)
// }
//
func Map(a interface{}, f func(a interface{}, i int) interface{}) (interface{}, error) {
	val := reflect.ValueOf(a)
	// valF := reflect.ValueOf(f)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var slice reflect.Value
		var kind reflect.Kind
		assign := false
		for i := 0; i < val.Len(); i++ {
			value := val.Index(i)
			result := f(value.Interface(), i)
			if !assign {
				t2 := reflect.ValueOf(result).Type()
				kind = t2.Kind()
				slice = reflect.Zero(reflect.SliceOf(t2))
				assign = true
			}
			slice = reflect.Append(slice, reflect.ValueOf(result))
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
		return slice.Slice, nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}
