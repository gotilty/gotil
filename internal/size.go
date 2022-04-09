package internal

import (
	"reflect"
	"strings"

	"github.com/gotilty/gotil/internal/errs"
)

func Size(val interface{}) (int, error) {
	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.Map, reflect.Struct, reflect.Array, reflect.Slice:
		return value.Len(), nil
	case reflect.String:
		val_str, err := ToString(val)
		if err == nil {
			return len(strings.Split(val_str, "")), nil
		}
		return 0, errs.CustomError(value.Kind().String())
	default:
		return 0, errs.CustomError(value.Kind().String())
	}
}
