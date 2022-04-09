package internal

import (
	"testing"
)

func TestIsNotAssigned(t *testing.T) {
	testData := getIsNotAssignedTestData()
	for key, test := range testData {
		a, b := test.output, IsNotAssigned(test.inputValue)
		if a != b {
			t.Errorf("IsNotAssigned does not works expected\ncase: %s\nexpected: %t taken: %t", key, a, b)
		}
	}
}

func BenchmarkIsNotAssignedString(b *testing.B) {
	testData := getIsNotAssignedTestData()
	for n := 0; n < b.N; n++ {
		IsNotAssigned(testData["string"].inputValue)
	}
}
func BenchmarkIsNotAssignedInteger(b *testing.B) {
	testData := getIsNotAssignedTestData()
	for n := 0; n < b.N; n++ {
		IsNotAssigned(testData["integer"].inputValue)
	}
}

func BenchmarkIsNotAssignedStruct(b *testing.B) {
	testData := getIsNotAssignedTestData()
	for n := 0; n < b.N; n++ {
		IsNotAssigned(testData["struct"].inputValue)
	}
}

type testStructIsNotAssigned struct {
	a int
}

func getIsNotAssignedTestData() map[string]struct {
	inputValue interface{}
	output     bool
} {
	_testStruct := &testStructIsNotAssigned{
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
			output:     false,
			// err:        errors.New("exception"),
		},
		"string1": {
			inputValue: "test string data",
			output:     false,
		},
		"float": {
			inputValue: 10.0,
			output:     false,
		},
		"empty_string": {
			inputValue: "",
			output:     true,
		},
		"nil_reference": {
			inputValue: nil,
			output:     true,
		},
		"struct": {
			inputValue: _testStruct,
			output:     false,
		},
		"string_array": {
			inputValue: stringArray,
			output:     false,
		},
	}
	return testData
}
