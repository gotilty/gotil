package converter_test

import (
	"errors"
	"testing"

	"github.com/gotilty/gotil/converter"
)

func TestConvertToInt8(t *testing.T) {
	testData := getConvertToInt8TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := converter.ToInt8(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToInt8 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToInt8String(b *testing.B) {
	testData := getConvertToInt8TestData()
	for n := 0; n < b.N; n++ {
		converter.ToInt8(testData["string"].inputValue)
	}
}
func BenchmarkConvertToInt8Integer(b *testing.B) {
	testData := getConvertToInt8TestData()
	for n := 0; n < b.N; n++ {
		converter.ToInt8(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToInt8Struct(b *testing.B) {
	testData := getConvertToInt8TestData()
	for n := 0; n < b.N; n++ {
		converter.ToInt8(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToInt8Uint(b *testing.B) {
	testData := getConvertToInt8TestData()
	for n := 0; n < b.N; n++ {
		converter.ToInt8(testData["uint"].inputValue)
	}
}

func getConvertToInt8TestData() map[string]struct {
	inputValue interface{}
	output     int8
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     int8
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     10,
			err:        nil,
		},
		"uint": {
			inputValue: uint64(18446744073709551615),
			output:     0,
			err:        errors.New("exception"),
		},
		"string1": {
			inputValue: "112",
			output:     112,
		},
		"float": {
			inputValue: 112.1254135,
			output:     112,
		},
		"empty_string": {
			inputValue: "",
			output:     0,
		},
		"nil_reference": {
			inputValue: nil,
			output:     0,
			err:        errors.New("exception"),
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
