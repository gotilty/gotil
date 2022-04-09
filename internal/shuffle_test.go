package internal

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestShuffle(t *testing.T) {
	testData := getShuffleTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := ShuffleSeed(test.inputValue, test.seed)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkShuffleIntegerSlice(b *testing.B) {
	testData := getShuffleTestData()
	for n := 0; n < b.N; n++ {
		Shuffle(testData["shuffle_numbers"].inputValue)
	}
}

func BenchmarkShuffleStructSlice(b *testing.B) {
	testData := getShuffleTestData()
	for n := 0; n < b.N; n++ {
		Shuffle(testData["shuffle_struct"].inputValue)
	}
}

func ExampleShuffle() {
	seed := time.Now().UnixNano()
	seed = 58239238
	//Seed you will get the same sequence of pseudo­random numbers
	// each time you run the program.

	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData, _ := ShuffleSeed(data, seed)
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
