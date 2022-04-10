package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSize(t *testing.T) {
	testData := getSizeTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := Size(test.inputValue)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkSizeIntegerSlice(b *testing.B) {
	testData := getSizeTestData()
	for n := 0; n < b.N; n++ {
		Size(testData["size_numbers"].inputValue)
	}
}

func BenchmarkSizeStructSlice(b *testing.B) {
	testData := getSizeTestData()
	for n := 0; n < b.N; n++ {
		Size(testData["size_struct"].inputValue)
	}
}

func ExampleSize() {
	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData, _ := Size(data)
	fmt.Println(newData)
	// Output: 4
}

func getSizeTestData() map[string]struct {
	inputValue interface{}
	output     int
	err        error
} {

	testData := map[string]struct {
		inputValue interface{}
		output     int
		err        error
	}{
		"size_numbers": {
			inputValue: []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55},
			output:     9,
			err:        nil,
		},
		"size_struct": {
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
			output: 4,
			err:    nil,
		},
		"size_string": {
			inputValue: "gotil is perfect!",
			output:     17,
			err:        nil,
		},
	}
	return testData
}
