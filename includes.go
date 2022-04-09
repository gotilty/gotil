package gotil

import "github.com/gotilty/gotil/internal"

func Includes(a interface{}, val interface{}) (bool, error) {
	return internal.Includes(a, val)
}
