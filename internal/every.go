package internal

import (
	"github.com/gotilty/gotil/internal/errs"
)

func Every(val, f interface{}) (bool, error) {
	state := true
	if err := Each(val, func(k, v interface{}) {
		if check, _ := Includes(v, f); !check {
			state = false
		}
	}); err == nil {
		return state, err
	} else {
		return false, errs.NewUnsupportedTypeError(err.Error())
	}
}
