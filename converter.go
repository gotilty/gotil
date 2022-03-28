package gotil

import (
	"reflect"

	"github.com/gotility/gotil/converter"
)

func ToInt32(a interface{}) (int32, error) {
	if a == nil {
		return 0, nil
	}
	val := reflect.ValueOf(a)
	return converter.ToInt32(val)
}
