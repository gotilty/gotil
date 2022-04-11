package internal

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

// FindLastBy iterates over elements of collection, returns an object of lastest element predicate if returns true for.
// 	data := []int64{-100, -5, 30, 100}
// 	newData, _ := FindLastBy(data, func(val interface{}, i int) bool {
// 		if val.(int64) == 30 {
// 			return true
// 		} else {
// 			return false
// 		}
// 	})
// 	// Output: 30
func FindLastBy(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	return FindLastByFromIndex(a, f, 0)
}

// FindLastBy iterates over elements of collection, returns an object of lastest element predicate if returns true for.
// It works same as FindLastBy just from parameter provides to set start index of given slice or array.
func FindLastByFromIndex(a interface{}, f func(val interface{}, i int) bool, from int) (interface{}, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		index := findLastIndex(val, f, 0)
		if index > -1 {
			return val.Index(index).Interface(), nil
		}
		return nil, nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}

func findLastIndex(val reflect.Value, f func(val interface{}, i int) bool, from int) int {
	length := val.Len()
	if length == 0 {
		return -1
	}
	index := length - 1
	if from < 0 {
		index = max(length+index, 0)
	} else {
		index = min(index, length-1)
	}
	return baseFindIndex(val, f, index, true)
}

func baseFindIndex(val reflect.Value, f func(val interface{}, i int) bool, from int, right bool) int {
	length := val.Len()
	index := from
	if right {
		for index < length {
			value := val.Index(index)
			if f(value.Interface(), index) {
				return index
			}
			index--
		}
	} else {
		for index < length {
			value := val.Index(index)
			if f(value.Interface(), index) {
				return index
			}
			index++
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
