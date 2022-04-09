package gotil

import (
	"github.com/gotilty/gotil/internal"
)

func GroupBy(a interface{}, f func(val, i interface{}) interface{}) (interface{}, error) {
	return internal.GroupBy(a, f)
}
