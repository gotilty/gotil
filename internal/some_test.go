package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSome(t *testing.T) {
	testData := getSomeTestData()
	for _, test := range testData {
		a, erra := test.output, test.err
		b, errb := Some(test.inputValue1, test.inputValue2)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf(errb.Error())
			}
		}
	}
}

func ExampleSome() {
	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	result, _ := Some(data, -100)
	fmt.Println(result)
	// Output: true
}

func getSomeTestData() map[string]struct {
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
			output:      true,
			err:         nil,
		},
		"every_int2": {
			inputValue1: []int64{-100, -100, -100},
			inputValue2: 100,
			output:      false,
			err:         nil,
		},
	}
	return testData
}
