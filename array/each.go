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
func Each(a interface{}, f func(a interface{}, i int)) error {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			value := val.Index(i)
			f(value.Interface(), i)
		}
		return nil
	}
	return errs.NewUnsupportedTypeError(val.Kind().String())
}
