package gotil

import (
	"github.com/gotilty/gotil/internal"
)

// Shuffle returns a new array of shuffled values, using a version of the Fisher-Yates shuffle.
//
// The default seed is time.Now().UnixNano()
//
// To change seed use gotil.ShuffleSeed()
//
// seed := time.Now().UnixNano()
// seed := int64(58239238)
// Seed you will get the same sequence of pseudo­random numbers
// each time you run the program.

// data := []int64{-100, -5, 30, 100}
// newData, _ := Shuffle(data)
// Output: [-5 100 -100 30]
func Shuffle(a interface{}) (interface{}, error) {
	return internal.Shuffle(a)
}

// ShuffleSeed returns a new array of shuffled values, using a version of the Fisher-Yates shuffle with given seed
//
// To use randomize seed -> gotil.ShuffleSeed()
//
// seed := time.Now().UnixNano()
// Seed you will get the same sequence of pseudo­random numbers
// each time you run the program.

// data := []int64{-100, -5, 30, 100}
// newData, _ := ShuffleSeed(data, seed)
// Output: [-5 100 -100 30]
func ShuffleSeed(a interface{}, seed int64) (interface{}, error) {
	return internal.ShuffleSeed(a, seed)

}
