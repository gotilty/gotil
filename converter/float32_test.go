package converter_test

import (
	"errors"
	"math"
	"testing"

	"github.com/gotilty/gotil/converter"
)

func TestConvertToFloat32(t *testing.T) {
	testData := getConvertToFloat32TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := converter.ToFloat32(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToFloat32 does not works expected\ncase: %s\nexpected: %f taken: %f error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToFloat32String(b *testing.B) {
	testData := getConvertToFloat32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToFloat32(testData["string"].inputValue)
	}
}
func BenchmarkConvertToFloat32Integer(b *testing.B) {
	testData := getConvertToFloat32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToFloat32(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToFloat32Struct(b *testing.B) {
	testData := getConvertToFloat32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToFloat32(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToFloat32Uint(b *testing.B) {
	testData := getConvertToFloat32TestData()
	for n := 0; n < b.N; n++ {
		converter.ToFloat32(testData["uint"].inputValue)
	}
}

func getConvertToFloat32TestData() map[string]struct {
	inputValue interface{}
	output     float32
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     float32
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     10,
			err:        nil,
		},
		"uint": {
			inputValue: math.MaxInt32,
			output:     float32(math.MaxInt32),
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
