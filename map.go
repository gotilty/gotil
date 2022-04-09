package gotil

import (
	"github.com/gotilty/gotil/internal"
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
	return internal.Map(a, f)
}
