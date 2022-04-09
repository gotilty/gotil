package converter_test

import (
	"errors"
	"math"
	"testing"

	"github.com/gotilty/gotil/converter"
)

func TestConvertToUint8(t *testing.T) {
	testData := getConvertToUint8TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := converter.ToUint8(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToUint8 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToUint8String(b *testing.B) {
	testData := getConvertToUint8TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint8(testData["string"].inputValue)
	}
}
func BenchmarkConvertToUint8Integer(b *testing.B) {
	testData := getConvertToUint8TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint8(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToUint8Struct(b *testing.B) {
	testData := getConvertToUint8TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint8(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToUint8Uint(b *testing.B) {
	testData := getConvertToUint8TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint8(testData["uint"].inputValue)
	}
}

func getConvertToUint8TestData() map[string]struct {
	inputValue interface{}
	output     uint8
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     uint8
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     10,
			err:        nil,
		},
		"uint": {
			inputValue: math.MaxInt8,
			output:     math.MaxInt8,
			err:        nil,
		},
		"string1": {
			inputValue: "255",
			output:     255,
		},
		"string2": {
			inputValue: "ertugrul",
			output:     0,
			err:        errors.New("exception"),
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
