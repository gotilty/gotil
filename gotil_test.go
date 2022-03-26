package gotil_test

import (
	"testing"

	"github.com/gotility/gotil"
)

type testStruct struct {
	a int
}

func TestIsAssigned(t *testing.T) {
	_test := &testStruct{
		a: 1,
	}
	tt := make([]string, 5)
	tests := map[string]struct {
		inputValue interface{}
		output     bool
		// err        error
	}{
		"successful conversion 1": {
			inputValue: 10,
			output:     true,
			// err:        errors.New("exception"),
		},
		"successful conversion 2": {
			inputValue: "test string data",
			output:     true,
		},
		"successful conversion 3": {
			inputValue: 10.0,
			output:     true,
		},
		"successful conversion 4": {
			inputValue: "",
			output:     false,
		},
		"successful conversion 5": {
			inputValue: nil,
			output:     false,
		},
		"successful conversion 6": {
			inputValue: _test,
			output:     true,
		},
		"successful conversion 7": {
			inputValue: tt,
			output:     true,
		},
	}
	for key, test := range tests {
		a, b := test.output, gotil.IsAssigned(test.inputValue)
		if a != b {
			t.Errorf("IsAssigned not works expected\ncase: %s\nexpected: %t taken: %t", key, a, b)
		}
	}
}
