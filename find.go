package gotil

import "github.com/gotilty/gotil/internal"

func FindBy(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	return FindBy(a, f)
}

func FindByFromIndex(a interface{}, f func(val interface{}, i int) bool, from int) (interface{}, error) {
	return internal.FindLastByFromIndex(a, f, from)
}
