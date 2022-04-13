package gotil

import (
	"github.com/gotilty/gotil/internal"
)

func Every(val, f interface{}) (bool, error) {
	return internal.Every(val, f)
}
