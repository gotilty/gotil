package gotil

import (
	"github.com/gotilty/gotil/internal"
)

func Some(val, f interface{}) (bool, error) {
	return internal.Some(val, f)
}
