package gotil

import "github.com/gotilty/gotil/internal"

// FilterBy iterates over elements of collection, returning an array of all elements predicate if returns true for.
//	newData, _ := FilterBy([]int64{-100, -5, 30, 100}, func(val interface{}, i int) bool {
//	if val.(int64) > 0 {
// 		return true
// 		} else {
// 			return false
// 		}
// 	})
// 	// Output: [30 100]
func FilterBy(a interface{}, f func(val interface{}, i int) bool) (interface{}, error) {
	return internal.FilterBy(a, f)
}
