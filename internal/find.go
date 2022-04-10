package internal

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

// FindBy iterates over elements of collection, returns an object of fist element predicate if returns true for.
//
// data := []int64{-100, -5, 30, 100}
//
// newData, _ := FindBy(data, func(val interface{}, i int) bool {
// 	if val.(int64) == 30 {
// 		return true
// 	} else {
// 		return false
// 	}
// })
// Output: 30
func FindBy(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	return FindByFromIndex(a, f, 0)
}

// FindByFromIndex iterates over elements of collection, returns an object of fist element predicate if returns true for.
// It works same as FindBy just from parameter provides to set start index of given slice or array.
func FindByFromIndex(a interface{}, f func(val interface{}, i int) bool, from int) (interface{}, error) {
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
