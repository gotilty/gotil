package collection

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
	"github.com/gotilty/gotil/operator"
)

func Reduce(a interface{}, f func(accumulator interface{}, val interface{}, i int) interface{}, accumulator interface{}) (interface{}, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		length := val.Len()
		index := -1
		newSlice := copySlice(a)
		if operator.IsAssigned(accumulator) && length > 0 {
			index++
			accumulator = newSlice.Index(index).Interface()
		}
		for index++; index < length; {
			rowVal := val.Index(index)
			accumulator = f(accumulator, rowVal.Interface(), index)
		}
		return accumulator, nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}
