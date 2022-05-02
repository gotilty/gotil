package gotil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

func TestMap(t *testing.T) {
	input := []int64{10, 20, 30}
	expected := []int64{100, 400, 900}
	result := gotil.Map(input, func(val int64, i int) int64 {
		return val * val
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FindLastBy does not works expected\ncase: %d\nexpected: %d taken: %d", input, expected, result)
	}
}

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gotil.Map([]int{15, 50}, func(v, i int) int {
			return v * v
		})
	}
}

func ExampleMap() {
	result := gotil.Map([]int{10, 20}, func(v, i int) int {
		return v * v
	})
	fmt.Println(result)
	// Output: 100, 400
}
