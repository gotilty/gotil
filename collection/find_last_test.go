package collection_test

import (
	"reflect"
	"testing"

	"github.com/gotilty/gotil/collection"
)

func TestFindLast(t *testing.T) {
	testData := getFindLastTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := collection.FindLastByPredicate(test.inputValue, test.mapFunction)
		if erra == nil {
			if !reflect.DeepEqual(a, b) {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v error: %s", key, a, b, errb.Error())
			}
		}
	}
}

// func ExampleFindByIterate() {
// 	data := []int64{-100, -5, 30, 100}
// 	// Input: [-100 -5 30 100]
// 	newData, _ := collection.FilterByIterate(data, func(val interface{}, i int) bool {
// 		if val.(int64) == 30 {
// 			return true
// 		} else {
// 			return false
// 		}
// 	})
// 	fmt.Println(newData)
// 	// Output: [30]
// }

func getFindLastTestData() map[string]struct {
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
		"find_last_number": {
			inputValue: []int64{10, 15, 100, 30, 50},
			mapFunction: func(val interface{}, i int) bool {
				if val.(int64) > 80 {
					return true
				} else {
					return false
				}
			},
			output: int64(100),
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
				{
					name: "Mike",
					age:  43,
				},
			},
			mapFunction: func(val interface{}, i int) bool {
				if val.(user).age == 42 {
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
