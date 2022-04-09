package internal

import (
	"errors"
	"testing"
)

func TestConvertToInt(t *testing.T) {
	testData := getConvertToIntTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := ToInt(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToInt does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToIntString(b *testing.B) {
	testData := getConvertToIntTestData()
	for n := 0; n < b.N; n++ {
		ToInt(testData["string"].inputValue)
	}
}
func BenchmarkConvertToIntInteger(b *testing.B) {
	testData := getConvertToIntTestData()
	for n := 0; n < b.N; n++ {
		ToInt(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToIntStruct(b *testing.B) {
	testData := getConvertToIntTestData()
	for n := 0; n < b.N; n++ {
		ToInt(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToIntUint(b *testing.B) {
	testData := getConvertToIntTestData()
	for n := 0; n < b.N; n++ {
		ToInt(testData["uint"].inputValue)
	}
}

func getConvertToIntTestData() map[string]struct {
	inputValue interface{}
	output     int
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     int
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
			inputValue: "532634573456577654",
			output:     532634573456577654,
		},
		"float": {
			inputValue: 111241242.1254135,
			output:     111241242,
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
