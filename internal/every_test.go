package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEvery(t *testing.T) {
	testData := getEveryTestData()
	for _, test := range testData {
		a, erra := test.output, test.err
		b, errb := Every(test.inputValue1, test.inputValue2)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf(errb.Error())
			}
		}
	}
}

func ExampleEvery() {
	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	result, _ := Every(data, -100)
	fmt.Println(result)
	// Output: false
}

func getEveryTestData() map[string]struct {
	inputValue1 interface{}
	inputValue2 interface{}
	output      bool
	err         error
} {

	testData := map[string]struct {
		inputValue1 interface{}
		inputValue2 interface{}
		output      bool
		err         error
	}{
		"every_int": {
			inputValue1: []int64{-100, -5, 30, 100},
			inputValue2: -100,
			output:      false,
			err:         nil,
		},
		"every_int2": {
			inputValue1: []int64{-100, -100, -100},
			inputValue2: -100,
			output:      true,
			err:         nil,
		},
	}
	return testData
}
