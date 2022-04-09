package crypto_test

import (
	"testing"

	"github.com/gotilty/gotil/crypto"
	"github.com/gotilty/gotil/internal/errs"
)

func TestMd5(t *testing.T) {
	testData := getMD5TestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := crypto.Md5(test.inputValue)
		if erra == nil {
			if a != b || errb != nil {
				t.Errorf("Convert.ToString does not works expected\ncase: %s\nexpected: %s taken: %s error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkMD5String(b *testing.B) {
	testData := getMD5TestData()
	for n := 0; n < b.N; n++ {
		crypto.Md5(testData["string"].inputValue)
	}
}

//TODO: implement benchmark for array or slice

func getMD5TestData() map[string]struct {
	inputValue interface{}
	output     string
	err        error
} {

	testData := map[string]struct {
		inputValue interface{}
		output     string
		err        error
	}{
		"integer": {
			inputValue: 10,
			output:     "d3d9446802a44259755d38e6d163e820",
		},
		"uint": {
			inputValue: uint64(18446744073709551615),
			output:     "7e7825a3d8588a756abd0ce2ed07e121",
		},
		"string": {
			inputValue: "These pretzels are making me thirsty.",
			output:     "b0804ec967f48520697662a204f5fe72",
		},
		"float": {
			inputValue: 11234550.1254135,
			output:     "35a88a78e54592530f436aacdf36adec",
		},
		"empty_string": {
			inputValue: "",
			output:     "",
			err:        errs.EmptyStringError(),
		},
		"nil_reference": {
			inputValue: nil,
			output:     "",
			err:        errs.NilReferenceTypeError(),
		},
	}
	return testData
}
