package internal

import (
	"fmt"
	"os"
	"testing"
)

func TestEach(t *testing.T) {
	testData := getEachTestData()
	for key, test := range testData {
		erra := test.err
		errb := Each(test.inputValue, test.mapFunction)
		if erra == nil {
			if errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase %s: error: %s", key, errb.Error())
			}
		}
	}
}

func ExampleEach() {
	data := []int64{10, 20, 30}
	_ = Each(data, func(k, v interface{}) {
		fmt.Printf("%d apples", v)
	})
	// Output: 10 apples20 apples30 apples
}
func printTestLine(i, val interface{}) {
	fmt.Fprintln(os.Stdout, val)
}

func getEachTestData() map[string]struct {
	inputValue  interface{}
	mapFunction func(k, v interface{})
	err         error
} {

	testData := map[string]struct {
		inputValue  interface{}
		mapFunction func(k, v interface{})
		err         error
	}{
		"string_slice": {
			inputValue:  []string{"gotilty", "gotil"},
			mapFunction: printTestLine,
			err:         nil,
		},
		"map": {
			inputValue:  map[int][]string{1: {"gotilty", "is", "perfect!"}, 2: {"gotil"}},
			mapFunction: printTestLine,
			err:         nil,
		},
	}
	return testData
}
