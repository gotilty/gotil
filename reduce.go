package gotil

import "github.com/gotilty/gotil/internal"

// Reduce iterates given collection and returns the accumulated result of running each element
// the last param is initial value of accumulator.
//
// data := []int{5, 10}
//
// result, _ := Reduce(data, func(accumulator, val interface{}, i int) interface{} {
// 	return accumulator.(int) + val.(int)
// }, 0)
//
// Output: 15
func Reduce(a interface{}, f func(accumulator interface{}, val interface{}, i int) interface{}, accumulator interface{}) (interface{}, error) {
	return internal.Reduce(a, f, accumulator)
}
