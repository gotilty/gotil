package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIncludes(t *testing.T) {
	testData := getIncludesTestData()
	for _, test := range testData {
		a, erra := test.output, test.err
		b, errb := Includes(test.inputValue, test.Includes)
		fmt.Println(a)
		fmt.Println(b)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Includes does not works expected\ncase")
			}
		}
	}
}

func BenchmarkIncludesMathFloor(b *testing.B) {
	testData := getIncludesTestData()
	for n := 0; n < b.N; n++ {
		Includes(testData["float_array"].inputValue, testData["float_array"].Includes)
	}
}

func BenchmarkIncludesStringLen(b *testing.B) {
	testData := getIncludesTestData()
	for n := 0; n < b.N; n++ {
		Includes(testData["map"].inputValue, testData["map"].Includes)
	}
}

func ExampleIncludes() {
	m1 := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}
	// Input: [5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001]
	result, _ := Includes(m1, "gotil")
	fmt.Println(result)
	// Output: false
}

func getIncludesTestData() map[string]struct {
	inputValue interface{}
	Includes   interface{}
	output     interface{}
	err        error
} {
	testData := map[string]struct {
		inputValue interface{}
		Includes   interface{}
		output     interface{}
		err        error
	}{
		"float_array": {
			inputValue: []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001},
			Includes:   10.5,
			output:     true,
			err:        nil,
		},
		"map": {
			inputValue: map[int][]string{1: {"gotilty"}, 2: {"gotil"}, 3: {"is"}, 4: {"perfect"}, 5: {"and"}, 6: {"easy", "to", "use"}},
			Includes:   "gotilty",
			output:     true,
			err:        nil,
		},
		"map2": {
			inputValue: map[int][]string{1: {"gotilty"}, 2: {"gotil"}, 3: {"is"}, 4: {"perfect"}, 5: {"and"}, 6: {"easy", "to", "use"}},
			Includes:   12,
			output:     false,
			err:        nil,
		},
	}
	return testData
}
