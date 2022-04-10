package gotil

import "github.com/gotilty/gotil/internal"

// It checks whether the given value exists in the given parameter.
//
// data := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}
//
// result, _ := Includes(data, "gotil")
//
// Output: false
func Includes(a interface{}, val interface{}) (bool, error) {
	return internal.Includes(a, val)
}
