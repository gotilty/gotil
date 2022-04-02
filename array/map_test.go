package array_test

// import (
// 	"errors"
// 	"math"
// 	"testing"

// 	"github.com/gotilty/gotil/array"
// 	"github.com/gotilty/gotil/converter"
// )

// func TestMap(t *testing.T) {
// 	testData := getMapTestData()
// 	for key, test := range testData {
// 		a, erra := test.output, test.err
// 		b, errb := array.Map(test.inputValue)
// 		if erra == nil {
// 			if a != b || errb != nil {
// 				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
// 			}
// 		}
// 	}
// }

// func BenchmarkMap(b *testing.B) {
// 	testData := getMapTestData()
// 	for n := 0; n < b.N; n++ {
// 		converter.ToUint64(testData["string"].inputValue)
// 	}
// }
// func BenchmarkConvertToUint64Integer(b *testing.B) {
// 	testData := getMapTestData()
// 	for n := 0; n < b.N; n++ {
// 		converter.ToUint64(testData["integer"].inputValue)
// 	}
// }

// func BenchmarkConvertToUint64Uint(b *testing.B) {
// 	testData := getMapTestData()
// 	for n := 0; n < b.N; n++ {
// 		converter.ToUint64(testData["uint"].inputValue)
// 	}
// }

// func square(a interface{}, i int) interface{} {
// 	b, _ := converter.ToUint64(a)
// 	return b * 2
// }

// func addChar(a interface{}, i int) interface{} {
// 	b, _ := converter.ToString(a)
// 	if i == 0 {
// 		return b + "->"
// 	}
// 	return b + " is the best utility tool for Golang!"
// }

// func getMapTestData() map[string]struct {
// 	inputValue interface{}
// 	output     interface{}
// 	err        error
// } {

// 	stringArray := make([]string, 5)
// 	testData := map[string]struct {
// 		inputValue interface{}
// 		inputValue func
// 		output     interface{}
// 		err        error
// 	}{
// 		"integer": {
// 			inputValue1: 10,
// 			inputValue2: square,
// 			output:      10,
// 			err:         nil,
// 		},
// 		"uint": {
// 			inputValue: math.MaxInt64,
// 			output:     math.MaxInt64,
// 			err:        nil,
// 		},
// 		"string1": {
// 			inputValue: "1215123123",
// 			output:     1215123123,
// 		},
// 		"string2": {
// 			inputValue: "ertugrul",
// 			output:     0,
// 			err:        errors.New("exception"),
// 		},
// 		"float": {
// 			inputValue: 11234550.1254135,
// 			output:     11234550,
// 		},
// 		"empty_string": {
// 			inputValue: "",
// 			output:     0,
// 		},
// 		"nil_reference": {
// 			inputValue: nil,
// 			output:     0,
// 		},
// 		"string_array": {
// 			inputValue: stringArray,
// 			output:     0,
// 			err:        errors.New("exception"),
// 		},
// 	}
// 	return testData
// }
