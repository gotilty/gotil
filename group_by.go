package gotil

import (
	"github.com/gotilty/gotil/internal"
)

// GroupBy iterates over elements of collection, returns an object of fist element predicate if returns true for.
//	data := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}
// 	if result, err := GroupBy(data, func(a interface{}, i interface{}) interface{} {
// 		return math.Floor(a.(float64))
// 	}); err == nil {
// 		fmt.Println(result)
// 	}
// 	// Output: map[4:[4.23] 5:[5 5.15 5.99] 10:[10.5 10 10.75] 20:[20] 100:[100 100.0001]]
func GroupBy(a interface{}, f func(val, i interface{}) interface{}) (interface{}, error) {
	return internal.GroupBy(a, f)
}
