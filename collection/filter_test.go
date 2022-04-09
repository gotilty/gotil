package collection_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil/collection"
)

type user struct {
	name     string
	age      int
	location location
}

type location struct {
	city string
}

func TestFilterBy(t *testing.T) {
	testData := getFilterByTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := collection.FilterBy(test.inputValue, test.mapFunction)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func ExampleFilterBy() {
	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData, _ := collection.FilterBy(data, func(val interface{}, i int) bool {
		if val.(int64) > 0 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(newData)
	// Output: [30 100]
}

func getFilterByTestData() map[string]struct {
	inputValue  interface{}
	mapFunction func(a interface{}, i int) bool
	output      interface{}
	err         error
} {

	testData := map[string]struct {
		inputValue  interface{}
		mapFunction func(a interface{}, i int) bool
		output      interface{}
		err         error
	}{
		"filter_positive_numbers": {
			inputValue: []int64{-100, -5, 30, 100},
			mapFunction: func(val interface{}, i int) bool {
				if val.(int64) > 0 {
					return true
				} else {
					return false
				}
			},
			output: []int64{30, 100},
			err:    nil,
		},
		"filter_by_name": {
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
			},
			mapFunction: func(val interface{}, i int) bool {
				if val.(user).name == "Olivia" {
					return true
				} else {
					return false
				}
			},
			output: []user{
				{
					name: "Olivia",
					age:  42,
				},
			},
			err: nil,
		},
	}
	return testData
}
