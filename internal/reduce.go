package internal

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

// Reduce iterates given collection and returns the accumulated result of running each element
// the last param is initial value of accumulator.
// 	data := []int{5, 10}
// 	result, _ := Reduce(data, func(accumulator, val interface{}, i int) interface{} {
// 		return accumulator.(int) + val.(int)
// 	}, 0)
// 	// Output: 15
func Reduce(a interface{}, f func(accumulator interface{}, val interface{}, i int) interface{}, accumulator interface{}) (interface{}, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		length := val.Len()
		index := -1
		newSlice := copySlice(a)
		if IsAssigned(accumulator) && length > 0 {
			index++
			accumulator = newSlice.Index(index).Interface()
		}
		for index++; index < length; {
			rowVal := val.Index(index)
			accumulator = f(accumulator, rowVal.Interface(), index)
		}
		return accumulator, nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}
