package gotil

import (
	"reflect"

	"github.com/gotilty/gotil/converter"
)

func ToInt32(a interface{}) (int32, error) {
	if a == nil {
		return 0, nil
	}
	val := reflect.ValueOf(a)
	return converter.ToInt32(val)
}

func ToString(a ...interface{}) string {
	if a == nil {
		return ""
	}
	val := reflect.ValueOf(a)
	return converter.ToString(val)
}
