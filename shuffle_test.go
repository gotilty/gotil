package gotil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

func TestShuffle(t *testing.T) {
	input := []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55}
	expected := []int64{100, 33, 1000, 30, 55, -5, -100, 5, 11}
	result := gotil.ShuffleSeed(input, int64(343434))
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FindLastBy does not works expected\ncase: %d\nexpected: %d taken: %d", input, expected, result)
	}
}

func BenchmarkShuffleIntegerSlice(b *testing.B) {
	input := []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55}
	for n := 0; n < b.N; n++ {
		gotil.Shuffle(input)
	}
}

func BenchmarkShuffleStructSlice(b *testing.B) {
	input := []user{
		{
			name: "Micheal",
			age:  27,
		},
		{
			name: "Joe",
			age:  30,
		},
		{
			name: "Olivia",
			age:  42,
		},
		{
			name: "Kevin",
			age:  10,
		},
	}
	for n := 0; n < b.N; n++ {
		gotil.Shuffle(input)
	}
}

func ExampleShuffle() {
	// seed := time.Now().UnixNano()
	seed := int64(58239238)
	//Seed you will get the same sequence of pseudo­random numbers
	// each time you run the program.

	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData := gotil.ShuffleSeed(data, seed)
	fmt.Println(newData)
	// Output: [-5 100 -100 30]
}

func getShuffleTestData() map[string]struct {
	inputValue interface{}
	output     interface{}
	seed       int64
	err        error
} {

	testData := map[string]struct {
		inputValue interface{}
		output     interface{}
		seed       int64
		err        error
	}{
		"shuffle_numbers": {
			inputValue: []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55},
			output:     []int64{100, 33, 1000, 30, 55, -5, -100, 5, 11},
			seed:       343434,
			err:        nil,
		},
		"shuffle_struct": {
			inputValue: []user{
				{
					name: "Micheal",
					age:  27,
				},
				{
					name: "Joe",
					age:  30,
				},
				{
					name: "Olivia",
					age:  42,
				},
				{
					name: "Kevin",
					age:  10,
				},
			},
			seed: 303030,
			output: []user{
				{
					name: "Olivia",
					age:  42,
				},
				{
					name: "Micheal",
					age:  27,
				},
				{
					name: "Kevin",
					age:  10,
				},
				{
					name: "Joe",
					age:  30,
				},
			},
			err: nil,
		},
	}
	return testData
}
