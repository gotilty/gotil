package internal

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/gotilty/gotil/config"
	"github.com/gotilty/gotil/internal/errs"
)

func TestToString(t *testing.T) {
	testData := getConvertToStringTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := ToString(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToString does not works expected\ncase: %s\nexpected: %s taken: %s error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkConvertToString_Int32(b *testing.B) {
	testData := getConvertToStringTestData()
	for n := 0; n < b.N; n++ {
		ToString(testData["integer"].inputValue)
	}
}

func BenchmarkConvertToString_Int32Array(b *testing.B) {
	testData := getConvertToStringTestData()
	for n := 0; n < b.N; n++ {
		ToString(testData["struct"].inputValue)
	}
}
func BenchmarkConvertToStringUInt(b *testing.B) {
	testData := getConvertToStringTestData()
	for n := 0; n < b.N; n++ {
		ToString(testData["uint"].inputValue)
	}
}

type testStruct struct {
	a int
}

func getConvertToStringTestData() map[string]struct {
	inputValue interface{}
	output     string
	err        error
} {
	_testStruct := &testStruct{
		a: 1,
	}
	arrayLenght := 3
	stringArray := make([]int, arrayLenght)
	var buffer bytes.Buffer
	for i := 0; i < arrayLenght; i++ {
		stringArray[i] = rand.Int()
		b, _ := ToString(stringArray[i])
		if i == arrayLenght-1 {
			buffer.WriteString(b)
		} else {
			buffer.WriteString(b + config.GetDefaultSeparator())
		}
	}
	stringArrayResult := buffer.String()

	testData := map[string]struct {
		inputValue interface{}
		output     string
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     "10",
		},
		"uint": {
			inputValue: uint64(18446744073709551615),
			output:     "18446744073709551615",
		},
		"string1": {
			inputValue: "1215123123",
			output:     "1215123123",
		},
		"float": {
			inputValue: 11234550.1254135,
			output:     "11234550.1254135",
		},
		"empty_string": {
			inputValue: "",
			output:     "",
		},
		"nil_reference": {
			inputValue: nil,
			output:     "",
			err:        errs.NilReferenceTypeError(),
		},
		"struct": {
			inputValue: _testStruct,
			output:     "",
			err:        errs.NewUnsupportedTypeError("struct"),
		},
		"string_array": {
			inputValue: stringArray,
			output:     stringArrayResult,
		},
	}
	return testData
}
