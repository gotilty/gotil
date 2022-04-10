package gotil

import (
	"github.com/gotilty/gotil/internal"
)

// FindLastBy iterates over elements of collection, returns an object of lastest element predicate if returns true for.
//
// data := []int64{-100, -5, 30, 100}
//
// newData, _ := FindLastBy(data, func(val interface{}, i int) bool {
// 	if val.(int64) == 30 {
// 		return true
// 	} else {
// 		return false
// 	}
// })
// Output: 30
func FindLastBy(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	return internal.FindLastBy(a, f)
}

// FindLastBy iterates over elements of collection, returns an object of lastest element predicate if returns true for.
// It works same as FindLastBy just from parameter provides to set start index of given slice or array.
func FindLastByFromIndex(a interface{}, f func(val interface{}, i int) bool, from int) (interface{}, error) {
	return internal.FindLastByFromIndex(a, f, from)
}
