package converter_test

import (
	"errors"
	"math"
	"testing"

	"github.com/gotilty/gotil/converter"
)

func TestConvertToUint32(t *testing.T) {
	testData := getConvertToUint32TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := converter.ToUint32(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToUint32 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToUint32String(b *testing.B) {
	testData := getConvertToUint32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint32(testData["string"].inputValue)
	}
}
func BenchmarkConvertToUint32Integer(b *testing.B) {
	testData := getConvertToUint32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint32(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToUint32Struct(b *testing.B) {
	testData := getConvertToUint32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint32(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToUint32Uint(b *testing.B) {
	testData := getConvertToUint32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToUint32(testData["uint"].inputValue)
	}
}

func getConvertToUint32TestData() map[string]struct {
	inputValue interface{}
	output     uint32
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     uint32
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     10,
			err:        nil,
		},
		"uint": {
			inputValue: math.MaxInt32,
			output:     math.MaxInt32,
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
