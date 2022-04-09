package gotil

import (
	"github.com/gotilty/gotil/internal"
)

func FindLastBy(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	return internal.FindLastBy(a, f)
}

func FindLastByFromIndex(a interface{}, f func(val interface{}, i int) bool, from int) (interface{}, error) {
	return internal.FindLastByFromIndex(a, f, from)
}
