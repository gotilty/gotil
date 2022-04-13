package internal

import (
	"errors"
	"math"
	"testing"
)

func TestConvertToUint64(t *testing.T) {
	testData := getConvertToUint64TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := ToUint64(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToUint64String(b *testing.B) {
	testData := getConvertToUint64TestData()
	for n := 0; n < b.N; n++ {
		ToUint64(testData["string"].inputValue)
	}
}

func BenchmarkConvertToUint64Integer(b *testing.B) {
	testData := getConvertToUint64TestData()
	for n := 0; n < b.N; n++ {
		ToUint64(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToUint64Struct(b *testing.B) {
	testData := getConvertToUint64TestData()
	for n := 0; n < b.N; n++ {
		ToUint64(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToUint64Uint(b *testing.B) {
	testData := getConvertToUint64TestData()
	for n := 0; n < b.N; n++ {
		ToUint64(testData["uint"].inputValue)
	}
}

func getConvertToUint64TestData() map[string]struct {
	inputValue interface{}
	output     uint64
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     uint64
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     10,
			err:        nil,
		},
		"uint": {
			inputValue: math.MaxInt64,
			output:     math.MaxInt64,
			err:        nil,
		},
		"string1": {
			inputValue: "1215123123",
			output:     1215123123,
		},
		"string2": {
			inputValue: "ertugrul",
			output:     0,
			err:        errors.New("exception"),
		},
		"float": {
			inputValue: 11234550.1254135,
			output:     11234550,
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
