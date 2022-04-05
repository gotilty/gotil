package array_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gotilty/gotil/array"
)

func TestEach(t *testing.T) {
	testData := getEachTestData()
	for key, test := range testData {
		erra := test.err
		errb := array.Each(test.inputValue, test.mapFunction)
		if erra == nil {
			if errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase %s: error: %s", key, errb.Error())
			}
		}
	}
}

func ExampleEach() {
	data := []int64{10, 20, 30}
	_ = array.Each(data, func(a interface{}, i int) {
		fmt.Printf("%d apples", a)
	})
	// Out: [10 apples 20 apples 30 apples]
	// Output: 10 apples20 apples30 apples
}
func printTestLine(a interface{}, i int) {

	fmt.Fprintln(os.Stdout, a)
}

func getEachTestData() map[string]struct {
	inputValue  interface{}
	mapFunction func(a interface{}, i int)
	err         error
} {

	testData := map[string]struct {
		inputValue  interface{}
		mapFunction func(a interface{}, i int)
		err         error
	}{
		"string": {
			inputValue:  []string{"gotilty", "gotil"},
			mapFunction: printTestLine,
			err:         nil,
		},
	}
	return testData
}
