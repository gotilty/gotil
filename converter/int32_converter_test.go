package converter_test

import (
	"errors"
	"testing"

	"github.com/gotilty/gotil/converter"
)

func TestConvertToInt32(t *testing.T) {
	testData := getConvertToInt32TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := converter.ToInt32(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToInt32 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToInt32String(b *testing.B) {
	testData := getConvertToInt32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToInt32(testData["string"].inputValue)
	}
}
func BenchmarkConvertToInt32Integer(b *testing.B) {
	testData := getConvertToInt32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToInt32(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToInt32Struct(b *testing.B) {
	testData := getConvertToInt32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToInt32(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToInt32Uint(b *testing.B) {
	testData := getConvertToInt32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToInt32(testData["uint"].inputValue)
	}
}

type testStruct struct {
	a int
}

func getConvertToInt32TestData() map[string]struct {
	inputValue interface{}
	output     int32
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     int32
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
			inputValue: "1215123123",
			output:     1215123123,
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
