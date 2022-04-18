package gotil

import (
	"math/rand"
	"time"
)

// Shuffle returns a new array of shuffled values, using a version of the Fisher-Yates shuffle.
// The default seed is time.Now().UnixNano()
// To change seed use gotil.ShuffleSeed()
// 	seed := int64(58239238)
//	Seed you will get the same sequence of pseudo­random numbers
// 	each time you run the program.
// 	data := []int64{-100, -5, 30, 100}
// 	newData := Shuffle(data)
// 	// Output: [-5 100 -100 30]
func Shuffle[T any](s []T) []T {
	return ShuffleSeed(s, time.Now().UnixNano())
}

// ShuffleSeed returns a new array of shuffled values, using a version of the Fisher-Yates shuffle with given seed
// To use randomize seed -> gotil.ShuffleSeed()
//
// seed := time.Now().UnixNano()
// Seed you will get the same sequence of pseudo­random numbers
// each time you run the program.
//
// 	data := []int64{-100, -5, 30, 100}
// 	newData := ShuffleSeed(data, seed)
// 	// Output: [-5 100 -100 30]
func ShuffleSeed[T any](s []T, seed int64) []T {
	rand.Seed(seed)
	newSlice := copySlice(s)
	for i := len(newSlice) - 1; i > 0; i-- { // Fisher–Yates shuffle
		j := rand.Intn(i + 1)
		newSlice[i], newSlice[j] = newSlice[j], newSlice[i]
	}
	return newSlice
}
