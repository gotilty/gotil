package gotil

import "github.com/gotilty/gotil/internal"

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
	return internal.FindBy(a, f)
}

// FindByFromIndex iterates over elements of collection, returns an object of fist element predicate if returns true for.
// It works same as FindBy just from parameter provides to set start index of given slice or array.
func FindByFromIndex(a interface{}, f func(val interface{}, i int) bool, from int) (interface{}, error) {
	return internal.FindByFromIndex(a, f, from)
}
