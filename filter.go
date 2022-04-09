package gotil

import "github.com/gotilty/gotil/internal"

func FilterBy(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	return internal.FilterBy(a, f)
}
