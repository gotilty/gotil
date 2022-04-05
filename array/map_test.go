package array_test

import (
	"reflect"
	"testing"

	"github.com/gotilty/gotil/array"
	"github.com/gotilty/gotil/converter"
)

func TestMap(t *testing.T) {
	testData := getMapTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := array.Map(test.inputValue, test.mapFunction)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkMapInt(b *testing.B) {
	testData := getMapTestData()
	for n := 0; n < b.N; n++ {
		array.Map(testData["integer2"].inputValue, testData["integer2"].mapFunction)
	}
}

func BenchmarkMapString(b *testing.B) {
	testData := getMapTestData()
	for n := 0; n < b.N; n++ {
		array.Map(testData["integer2"].inputValue, testData["integer2"].mapFunction)
	}
}

func BenchmarkIntToString(b *testing.B) {
	testData := getMapTestData()
	for n := 0; n < b.N; n++ {
		array.Map(testData["integer2"].inputValue, testData["integer2"].mapFunction)
	}
}

func square(a interface{}, i int) interface{} {
	return a.(int64) * a.(int64)
}

func addChar(a interface{}, i int) interface{} {
	b, _ := converter.ToString(a)
	if i == 0 {
		return b
	}
	return b + " is perfect"
}

func intToString(a interface{}, i int) interface{} {
	b, _ := converter.ToString(a)
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
