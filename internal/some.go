package internal

import (
	"github.com/gotilty/gotil/internal/errs"
)

func Some(val, f interface{}) (bool, error) {
	state := false
	if err := Each(val, func(k, v interface{}) {
		if check, _ := Includes(v, f); !check {
			state = true
		}
	}); err == nil {
		return state, err
	} else {
		return false, errs.NewUnsupportedTypeError(err.Error())
	}
}
