package collection_test

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/gotilty/gotil/collection"
	"github.com/gotilty/gotil/converter"
)

func TestGroupBy(t *testing.T) {
	testData := getGroupByTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := collection.GroupBy(test.inputValue, test.groupByFunction)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("collection.GroupBy does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func BenchmarkGroupByMathFloor(b *testing.B) {
	testData := getGroupByTestData()
	for n := 0; n < b.N; n++ {
		collection.GroupBy(testData["math_floor"].inputValue, testData["math_floor"].groupByFunction)
	}
}

func BenchmarkGroupByStringLen(b *testing.B) {
	testData := getGroupByTestData()
	for n := 0; n < b.N; n++ {
		collection.GroupBy(testData["string_len"].inputValue, testData["string_len"].groupByFunction)
	}
}

func ExampleGroupBy() {
	m1 := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}
	// Input: [5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001]
	if result, err := collection.GroupBy(m1, func(a interface{}, i interface{}) interface{} {
		return math.Floor(a.(float64))
	}); err == nil {
		fmt.Println(result)
	}
	// Output: map[4:[4.23] 5:[5 5.15 5.99] 10:[10.5 10 10.75] 20:[20] 100:[100 100.0001]]
}

func getGroupByTestData() map[string]struct {
	inputValue      interface{}
	groupByFunction func(val interface{}, i interface{}) interface{}
	output          interface{}
	err             error
} {
	testData := map[string]struct {
		inputValue      interface{}
		groupByFunction func(val interface{}, i interface{}) interface{}
		output          interface{}
		err             error
	}{
		"math_floor": {
			inputValue: []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001},
			groupByFunction: func(val interface{}, i interface{}) interface{} {
				return math.Floor(val.(float64))
			},
			output: map[float64][]float64{4: {4.23}, 5: {5, 5.15, 5.99}, 10: {10.5, 10, 10.75}, 20: {20}, 100: {100, 100.0001}},
			err:    nil,
		},
		"string_len": {
			inputValue: []string{"gotilty", "gotil", "is", "perfect", "and", "easy", "to", "use"},
			groupByFunction: func(val, i interface{}) interface{} {
				value, _ := converter.ToString(val)
				return len(strings.Split(value, ""))
			},
			output: map[int][]string{2: {"is", "to"}, 3: {"and", "use"}, 4: {"easy"}, 5: {"gotil"}, 7: {"gotilty", "perfect"}},
			err:    nil,
		},
		// "map_groupby": {
		// 	inputValue: map[int][]string{1: {"gotilty"}, 2: {"gotil"}, 3: {"is"}, 4: {"perfect"}, 5: {"and"}, 6: {"easy"}, 7: {"to"}, 8: {"use"}},
		// 	groupByFunction: func(val, i interface{}) interface{} {
		// 		for j := 0; j < len(val.([]string)); j++ {
		// 			value, _ := converter.ToString(val)
		// 			return len(strings.Split(value, ""))
		// 		}
		// 		return len(strings.Split(val.(string), ""))
		// 	},
		// 	output: map[int][]string{2: {"is", "to"}, 3: {"and", "use"}, 4: {"easy"}, 5: {"gotil"}, 7: {"gotilty", "perfect"}},
		// 	err:    nil,
		// },
	}
	return testData
}
