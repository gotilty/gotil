package converter_test

import (
	"errors"
	"math"
	"testing"

	"github.com/gotilty/gotil/converter"
)

func TestConvertToUint16(t *testing.T) {
	testData := getConvertToUint16TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := converter.ToUint16(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToUint16 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToUint16String(b *testing.B) {
	testData := getConvertToUint16TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint16(testData["string"].inputValue)
	}
}
func BenchmarkConvertToUint16Integer(b *testing.B) {
	testData := getConvertToUint16TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint16(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToUint16Struct(b *testing.B) {
	testData := getConvertToUint16TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint16(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToUint16Uint(b *testing.B) {
	testData := getConvertToUint16TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint16(testData["uint"].inputValue)
	}
}

func getConvertToUint16TestData() map[string]struct {
	inputValue interface{}
	output     uint16
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     uint16
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     10,
			err:        nil,
		},
		"uint": {
			inputValue: math.MaxInt16,
			output:     math.MaxInt16,
			err:        nil,
		},
		"string1": {
			inputValue: "65535",
			output:     65535,
		},
		"string2": {
			inputValue: "ertugrul",
			output:     0,
			err:        errors.New("exception"),
		},
		"float": {
			inputValue: 11234.1254135,
			output:     11234,
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
