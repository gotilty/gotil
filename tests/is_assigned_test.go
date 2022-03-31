package tests

import (
	"testing"

	"github.com/gotilty/gotil"
)

func TestIsAssigned(t *testing.T) {
	testData := getIsAssignedTestData()
	for key, test := range testData {
		a, b := test.output, gotil.IsAssigned(test.inputValue)
		if a != b {
			t.Errorf("IsAssigned does not works expected\ncase: %s\nexpected: %t taken: %t", key, a, b)
		}
	}
}

func BenchmarkIsAssignedString(b *testing.B) {
	testData := getIsAssignedTestData()
	for n := 0; n < b.N; n++ {
		gotil.IsAssigned(testData["string"].inputValue)
	}
}
func BenchmarkIsAssignedInteger(b *testing.B) {
	testData := getIsAssignedTestData()
	for n := 0; n < b.N; n++ {
		gotil.IsAssigned(testData["integer"].inputValue)
	}
}

func BenchmarkIsAssignedStruct(b *testing.B) {
	testData := getIsAssignedTestData()
	for n := 0; n < b.N; n++ {
		gotil.IsAssigned(testData["struct"].inputValue)
	}
}

type testStruct struct {
	a int
}

func getIsAssignedTestData() map[string]struct {
	inputValue interface{}
	output     bool
} {
	_testStruct := &testStruct{
		a: 1,
	}
	stringArray := make([]string, 5)
	testData := map[string]struct {
		inputValue interface{}
		output     bool
		// err        error
	}{
		"integer": {
			inputValue: 10,
			output:     true,
			// err:        errors.New("exception"),
		},
		"string1": {
			inputValue: "test string data",
			output:     true,
		},
		"float": {
			inputValue: 10.0,
			output:     true,
		},
		"empty_string": {
			inputValue: "",
			output:     false,
		},
		"nil_reference": {
			inputValue: nil,
			output:     false,
		},
		"struct": {
			inputValue: _testStruct,
			output:     true,
		},
		"string_array": {
			inputValue: stringArray,
			output:     true,
		},
	}
	return testData
}
