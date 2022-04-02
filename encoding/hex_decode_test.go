package encoding_test

import (
	"testing"

	"github.com/gotilty/gotil/converter"
	"github.com/gotilty/gotil/encoding"
)

func TestHexDecode(t *testing.T) {
	testData := getHexDecodeTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := encoding.HexDecode(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Encoding.HexDecode does not works expected\ncase: %s\nexpected: %s taken: %s ", key, a, b)
			}
		}
	}
}

func BenchmarkConvertToHexToString(b *testing.B) {
	testData := getHexDecodeTestData()
	for n := 0; n < b.N; n++ {
		converter.ToString(testData["string"].inputValue)
	}
}

// func BenchmarkConvertToString_Int32Array(b *testing.B) {
// 	testData := getHexEncodeTestData()
// 	for n := 0; n < b.N; n++ {
// 		converter.ToString(testData["struct"].inputValue)
// 	}
// }
// func BenchmarkConvertToStringUInt(b *testing.B) {
// 	testData := getHexEncodeTestData()
// 	for n := 0; n < b.N; n++ {
// 		converter.ToString(testData["uint"].inputValue)
// 	}
// }

func getHexDecodeTestData() map[string]struct {
	inputValue interface{}
	output     string
	err        error
} {
	// stringArrayResult := buffer.String()

	testData := map[string]struct {
		inputValue interface{}
		output     string
		err        error
	}{
		"string": {
			inputValue: "48656c6c6f20476f7068657221",
			output:     "Hello Gopher!",
			err:        nil,
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
