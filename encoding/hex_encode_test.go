package encoding_test

import (
	"bytes"
	"math/rand"
	"testing"

	"github.com/gotilty/gotil/config"
	"github.com/gotilty/gotil/converter"
	"github.com/gotilty/gotil/encoding"
)

func TestToString(t *testing.T) {
	testData := getConvertToStringTestData()
	for key, test := range testData {
		a := test.output
		b := encoding.HexEncode(test.inputValue)

		if a != b {
			t.Errorf("Encoding.HexEncode does not works expected\ncase: %s\nexpected: %s taken: %s ", key, a, b)
		}
	}
}

func BenchmarkConvertToString_Int32(b *testing.B) {
	testData := getConvertToStringTestData()
	for n := 0; n < b.N; n++ {
		converter.ToString(testData["string"].inputValue)
	}
}

// func BenchmarkConvertToString_Int32Array(b *testing.B) {
// 	testData := getConvertToStringTestData()
// 	for n := 0; n < b.N; n++ {
// 		converter.ToString(testData["struct"].inputValue)
// 	}
// }
// func BenchmarkConvertToStringUInt(b *testing.B) {
// 	testData := getConvertToStringTestData()
// 	for n := 0; n < b.N; n++ {
// 		converter.ToString(testData["uint"].inputValue)
// 	}
// }

type testStruct struct {
	a int
}

func getConvertToStringTestData() map[string]struct {
	inputValue interface{}
	output     string
} {
	// _testStruct := &testStruct{
	// 	a: 1,
	// }
	arrayLenght := 3
	stringArray := make([]int, arrayLenght)
	var buffer bytes.Buffer
	for i := 0; i < arrayLenght; i++ {
		stringArray[i] = rand.Int()
		b := converter.ToString(stringArray[i])
		if i == arrayLenght-1 {
			buffer.WriteString(b)
		} else {
			buffer.WriteString(b + config.GetDefaultSeperator())
		}
	}
	// stringArrayResult := buffer.String()

	testData := map[string]struct {
		inputValue interface{}
		output     string
	}{
		"string": {
			inputValue: "Hello Gopher!",
			output:     "48656c6c6f20476f7068657221",
		},
		// "uint": {
		// 	inputValue: uint64(18446744073709551615),
		// 	output:     "18446744073709551615",
		// },
		// "string1": {
		// 	inputValue: "1215123123",
		// 	output:     "1215123123",
		// },
		// "float": {
		// 	inputValue: 11234550.1254135,
		// 	output:     "11234550.1254135",
		// },
		// "empty_string": {
		// 	inputValue: "",
		// 	output:     "",
		// },
		// "nil_reference": {
		// 	inputValue: nil,
		// 	output:     "",
		// },
		// "struct": {
		// 	inputValue: _testStruct,
		// 	output:     "",
		// },
		// "string_array": {
		// 	inputValue: stringArray,
		// 	output:     stringArrayResult,
		// },
	}
	return testData
}
