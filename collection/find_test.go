package collection_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil/collection"
)

func TestFindBy(t *testing.T) {
	testData := getFindByTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := collection.FindBy(test.inputValue, test.mapFunction)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkFindByIntegerSlice(b *testing.B) {
	testData := getFindByTestData()
	for n := 0; n < b.N; n++ {
		collection.FindBy(testData["find_number"].inputValue, testData["find_number"].mapFunction)
	}
}

func BenchmarkFindByStructSlice(b *testing.B) {
	testData := getFindByTestData()
	for n := 0; n < b.N; n++ {
		collection.FindBy(testData["find_by_name"].inputValue, testData["find_by_name"].mapFunction)
	}
}

func ExampleFindBy() {
	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData, _ := collection.FindBy(data, func(val interface{}, i int) bool {
		if val.(int64) == 30 {
			return true
		} else {
			return false
		}
	})
	fmt.Println(newData)
	// Output: 30
}

func getFindByTestData() map[string]struct {
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
		"find_number": {
			inputValue: []int64{-100, -5, 30, 100},
			mapFunction: func(val interface{}, i int) bool {
				if val.(int64) == 30 {
					return true
				} else {
					return false
				}
			},
			output: int64(30),
			err:    nil,
		},
		"find_by_name": {
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
			output: user{
				name: "Olivia",
				age:  42,
			},
			err: nil,
		},
	}
	return testData
}
