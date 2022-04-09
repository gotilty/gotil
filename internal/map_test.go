package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	testData := getMapTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := Map(test.inputValue, test.mapFunction)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func ExampleMap() {
	data := []int64{10, 20, 30}
	newData, _ := Map(data, func(val interface{}, i int) interface{} {
		return fmt.Sprintf("%d apples", val)
	})
	fmt.Println(newData)
	// Output: [10 apples 20 apples 30 apples]
}

func BenchmarkMapInt(b *testing.B) {
	testData := getMapTestData()
	for n := 0; n < b.N; n++ {
		Map(testData["integer2"].inputValue, testData["integer2"].mapFunction)
	}
}

func BenchmarkMapString(b *testing.B) {
	testData := getMapTestData()
	for n := 0; n < b.N; n++ {
		Map(testData["integer2"].inputValue, testData["integer2"].mapFunction)
	}
}

func BenchmarkIntToString(b *testing.B) {
	testData := getMapTestData()
	for n := 0; n < b.N; n++ {
		Map(testData["integer2"].inputValue, testData["integer2"].mapFunction)
	}
}

func square(val interface{}, i int) interface{} {
	return val.(int64) * val.(int64)
}

func addChar(val interface{}, i int) interface{} {
	b, _ := ToString(val)
	if i == 0 {
		return b
	}
	return b + " is perfect"
}

func intToString(val interface{}, i int) interface{} {
	b, _ := ToString(val)
	if i == 0 {
		return "gotilty " + b
	} else {
		return "gotil " + b
	}
}

func getMapTestData() map[string]struct {
	inputValue  interface{}
	mapFunction func(a interface{}, i int) interface{}
	output      interface{}
	err         error
} {

	testData := map[string]struct {
		inputValue  interface{}
		mapFunction func(a interface{}, i int) interface{}
		output      interface{}
		err         error
	}{
		"integer": {
			inputValue:  []int64{10, 20, 30},
			mapFunction: square,
			output:      []int64{100, 400, 900},
			err:         nil,
		},
		"string": {
			inputValue:  []string{"gotilty", "gotil"},
			mapFunction: addChar,
			output:      []string{"gotilty", "gotil is perfect"},
			err:         nil,
		}, "integer2": {
			inputValue:  []int{5, 10},
			mapFunction: intToString,
			output:      []string{"gotilty 5", "gotil 10"},
			err:         nil,
		},
	}
	return testData
}
