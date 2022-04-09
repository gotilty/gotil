package internal

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

// Creates an array of values by running each element in the given array thru iteratee.
// The value to be iterated should be given as the first parameter.
// The second parameter will be a function that will take the parameter and return value type interface{}.
// Example: func square(a interface{}, i int()) interface{} {
// 	b, _ := ToInt64(a)
// 	return (b * b)
// }
//
func Map(a interface{}, f func(val interface{}, i int) interface{}) (interface{}, error) {
	val := reflect.ValueOf(a)
	// valF := reflect.ValueOf(f)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var slice reflect.Value
		assign := false
		for i := 0; i < val.Len(); i++ {
			value := val.Index(i)
			result := f(value.Interface(), i)
			if !assign {
				t2 := reflect.ValueOf(result).Type()
				slice = reflect.Zero(reflect.SliceOf(t2))
				assign = true
			}
			slice = reflect.Append(slice, reflect.ValueOf(result))
		}
		return slice.Interface(), nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}
