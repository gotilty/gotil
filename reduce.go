package gotil

import "github.com/gotilty/gotil/internal"

func Reduce(a interface{}, f func(accumulator interface{}, val interface{}, i int) interface{}, accumulator interface{}) (interface{}, error) {
	return internal.Reduce(a, f, accumulator)
}
