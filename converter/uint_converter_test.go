package converter_test

import (
	"errors"
	"testing"

	"github.com/gotilty/gotil/converter"
)

func TestConvertToUint(t *testing.T) {
	testData := getConvertToUintTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := converter.ToUint(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToUint does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToUintString(b *testing.B) {
	testData := getConvertToUintTestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint(testData["string"].inputValue)
	}
}
func BenchmarkConvertToUintInteger(b *testing.B) {
	testData := getConvertToUintTestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToUintStruct(b *testing.B) {
	testData := getConvertToUintTestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToUintUint(b *testing.B) {
	testData := getConvertToUintTestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint(testData["uint"].inputValue)
	}
}

func getConvertToUintTestData() map[string]struct {
	inputValue interface{}
	output     uint
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     uint
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     10,
			err:        nil,
		},
		"uint": {
			inputValue: 1844674407370955161,
			output:     1844674407370955161,
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
