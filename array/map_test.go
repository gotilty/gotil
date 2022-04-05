package array_test

import (
	"testing"

	"github.com/gotilty/gotil/array"
	"github.com/gotilty/gotil/converter"
)

func TestMap(t *testing.T) {
	testData := getMapTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := array.Map(testData["string"].inputValue1, testData["string"].mapCallbackMethod)
		if erra == nil {
			if a == b || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkMap(b *testing.B) {
	testData := getMapTestData()
	for n := 0; n < b.N; n++ {
		array.Map(testData["string"].inputValue1, testData["string"].mapCallbackMethod)
	}
}

func square(a interface{}, i int) interface{} {
	b, _ := converter.ToUint64(a)
	return b * b
}

func addChar(a interface{}, i int) interface{} {
	b, _ := converter.ToString(a)
	if i == 0 {
		return b + "->"
	}
	return b + " is the best utility tool for Golang!"
}

func getMapTestData() map[string]struct {
	inputValue1       interface{}
	mapCallbackMethod func(i interface{}, a int) interface{}
	output            interface{}
	err               error
} {

	testData := map[string]struct {
		inputValue1       interface{}
		mapCallbackMethod func(i interface{}, a int) interface{}
		output            interface{}
		err               error
	}{
		"integer": {
			inputValue1:       []int{10, 20, 30},
			mapCallbackMethod: square,
			output:            []int{10, 20, 30},
			err:               nil,
		},
		// "uint": {
		// 	inputValue2: math.MaxInt64,
		// 	output:     math.MaxInt64,
		// 	err:        nil,
		// },
		// "string1": {
		// 	inputValue: "1215123123",
		// 	output:     1215123123,
		// },
		// "string2": {
		// 	inputValue: "ertugrul",
		// 	output:     0,
		// 	err:        errors.New("exception"),
		// },
		// "float": {
		// 	inputValue: 11234550.1254135,
		// 	output:     11234550,
		// },
		// "empty_string": {
		// 	inputValue: "",
		// 	output:     0,
		// },
		// "nil_reference": {
		// 	inputValue: nil,
		// 	output:     0,
		// },
		// "string_array": {
		// 	inputValue: stringArray,
		// 	output:     0,
		// 	err:        errors.New("exception"),
		// },
	}
	return testData
}
