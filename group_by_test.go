package gotil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

func TestGroupBy(t *testing.T) {
	input := []string{"gotilty", "gotil", "is", "perfect", "and", "easy", "to", "use"}
	expected := map[int][]string{2: {"is", "to"}, 3: {"and", "use"}, 4: {"easy"}, 5: {"gotil"}, 7: {"gotilty", "perfect"}}
	result := gotil.GroupBy(input, func(v string, i int) int {
		return len(v)
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("TestGroupBy does not works expected\ncase: %v\nexpected: %v taken: %v", input, expected, result)
	}
}

func BenchmarkGroupBy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = gotil.GroupBy([]int{10, 20, 30}, func(v int, i int) int {
			return v / 10
		})
	}
}

func ExampleGroupBy() {
	input := []int{10, 20, 30}
	result := gotil.GroupBy(input, func(v int, i int) int {
		return v / 10
	})
	fmt.Println(result)
	// Output: map[1:[10] 2:[20] 3:[30]]
}

// func BenchmarkGroupByMathFloor(b *testing.B) {
// 	testData := getGroupByTestData()
// 	for n := 0; n < b.N; n++ {
// 		GroupBy(testData["math_floor"].inputValue, testData["math_floor"].groupByFunction)
// 	}
// }

// func BenchmarkGroupByStringLen(b *testing.B) {
// 	testData := getGroupByTestData()
// 	for n := 0; n < b.N; n++ {
// 		GroupBy(testData["string_len"].inputValue, testData["string_len"].groupByFunction)
// 	}
// }

// func ExampleGroupBy() {
// 	data := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}
// 	// Input: [5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001]
// 	if result, err := GroupBy(data, func(a interface{}, i interface{}) interface{} {
// 		return math.Floor(a.(float64))
// 	}); err == nil {
// 		fmt.Println(result)
// 	}
// 	// Output: map[4:[4.23] 5:[5 5.15 5.99] 10:[10.5 10 10.75] 20:[20] 100:[100 100.0001]]
// }

// // func getGroupByTestData() map[string]struct {
// // 	inputValue      interface{}
// // 	groupByFunction func(val interface{}, i interface{}) interface{}
// 	output          interface{}
// 	err             error
// } {
// 	testData := map[string]struct {
// 		inputValue      interface{}
// 		groupByFunction func(val interface{}, i interface{}) interface{}
// 		output          interface{}
// 		err             error
// 	}{
// 		"math_floor": {
// 			inputValue: []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001},
// 			groupByFunction: func(val interface{}, i interface{}) interface{} {
// 				return math.Floor(val.(float64))
// 			},
// 			output: map[float64][]float64{4: {4.23}, 5: {5, 5.15, 5.99}, 10: {10.5, 10, 10.75}, 20: {20}, 100: {100, 100.0001}},
// 			err:    nil,
// 		},
// 		"string_len": {
// 			inputValue: []string{"gotilty", "gotil", "is", "perfect", "and", "easy", "to", "use"},
// 			groupByFunction: func(val, i interface{}) interface{} {
// 				value, _ := ToString(val)
// 				return len(strings.Split(value, ""))
// 			},
// 			output: map[int][]string{2: {"is", "to"}, 3: {"and", "use"}, 4: {"easy"}, 5: {"gotil"}, 7: {"gotilty", "perfect"}},
// 			err:    nil,
// 		},
// 		// "map_groupby": {
// 		// 	inputValue: map[int][]string{1: {"gotilty"}, 2: {"gotil"}, 3: {"is"}, 4: {"perfect"}, 5: {"and"}, 6: {"easy"}, 7: {"to"}, 8: {"use"}},
// 		// 	groupByFunction: func(val, i interface{}) interface{} {
// 		// 		for j := 0; j < len(val.([]string)); j++ {
// 		// 			value, _ := ToString(val)
// 		// 			return len(strings.Split(value, ""))
// 		// 		}
// 		// 		return len(strings.Split(val.(string), ""))
// 		// 	},
// 		// 	output: map[int][]string{2: {"is", "to"}, 3: {"and", "use"}, 4: {"easy"}, 5: {"gotil"}, 7: {"gotilty", "perfect"}},
// 		// 	err:    nil,
// 		// },
// 	}
// 	return testData
// }
