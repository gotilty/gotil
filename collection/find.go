package collection

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

func FindByPredicate(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	return FindByPredicateFromIndex(a, f, 0)

}

func FindByPredicateFromIndex(a interface{}, f func(val interface{}, i int) bool, from int) (interface{}, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		index := findIndex(val, f, 0)
		if index > -1 {
			return val.Index(index).Interface(), nil
		}
		return nil, nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}

func findIndex(val reflect.Value, f func(val interface{}, i int) bool, from int) int {
	length := val.Len()
	if length == 0 {
		return -1
	}
	index := 0
	if from < 0 {
		index = max(length+index, 0)
	} else {
		index = min(index, length-1)
	}
	return baseFindIndex(val, f, index, false)
}
