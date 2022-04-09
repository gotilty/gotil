package internal

import (
	"errors"
	"math"
	"testing"
)

func TestConvertToFloat64(t *testing.T) {
	testData := getConvertToFloat64TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := ToFloat64(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToFloat64 does not works expected\ncase: %s\nexpected: %f taken: %f error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToFloat64String(b *testing.B) {
	testData := getConvertToFloat64TestData()
	for n := 0; n < b.N; n++ {
		ToFloat64(testData["string"].inputValue)
	}
}
func BenchmarkConvertToFloat64Integer(b *testing.B) {
	testData := getConvertToFloat64TestData()
	for n := 0; n < b.N; n++ {
		ToFloat64(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToFloat64Struct(b *testing.B) {
	testData := getConvertToFloat64TestData()
	for n := 0; n < b.N; n++ {
		ToFloat64(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToFloat64Uint(b *testing.B) {
	testData := getConvertToFloat64TestData()
	for n := 0; n < b.N; n++ {
		ToFloat64(testData["uint"].inputValue)
	}
}

func getConvertToFloat64TestData() map[string]struct {
	inputValue interface{}
	output     float64
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     float64
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     10,
			err:        nil,
		},
		"uint": {
			inputValue: math.MaxInt64,
			output:     float64(math.MaxInt64),
			err:        nil,
		},
		"string1": {
			inputValue: "11234550.1254135",
			output:     11234550.1254135,
		},
		"string2": {
			inputValue: "ertugrul",
			output:     0,
			err:        errors.New("exception"),
		},
		"float": {
			inputValue: 11234550.1254135,
			output:     11234550.1254135,
		},
		"empty_string": {
			inputValue: "",
			output:     0,
		},
		"nil_reference": {
			inputValue: nil,
			output:     0,
		},
		"struct": {
			inputValue: _testStruct,
			output:     0,
			err:        errors.New("exception"),
		},
		"string_array": {
			inputValue: stringArray,
			output:     0,
			err:        errors.New("exception"),
		},
	}
	return testData
}
