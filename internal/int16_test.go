package internal

import (
	"errors"
	"testing"
)

func TestConvertToInt16(t *testing.T) {
	testData := getConvertToInt16TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := ToInt16(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToInt16 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToInt16String(b *testing.B) {
	testData := getConvertToInt16TestData()
	for n := 0; n < b.N; n++ {
		ToInt16(testData["string"].inputValue)
	}
}
func BenchmarkConvertToInt16Integer(b *testing.B) {
	testData := getConvertToInt16TestData()
	for n := 0; n < b.N; n++ {
		ToInt16(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToInt16Struct(b *testing.B) {
	testData := getConvertToInt16TestData()
	for n := 0; n < b.N; n++ {
		ToInt16(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToInt16Uint(b *testing.B) {
	testData := getConvertToInt16TestData()
	for n := 0; n < b.N; n++ {
		ToInt16(testData["uint"].inputValue)
	}
}

func getConvertToInt16TestData() map[string]struct {
	inputValue interface{}
	output     int16
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     int16
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
			inputValue: "12412",
			output:     12412,
		},
		"float": {
			inputValue: 11230.1254135,
			output:     11230,
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
