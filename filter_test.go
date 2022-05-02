package gotil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

func TestFilterBy(t *testing.T) {
	input := []user{
		{
			name: "Micheal",
			age:  27,
		},
		{
			name: "Joe",
			age:  30,
		},
		{
			name: "Olivia",
			age:  42,
		},
	}
	expected := []user{
		{
			name: "Olivia",
			age:  42,
		},
	}
	result := gotil.FilterBy(input, func(v user, i int) bool {
		return v.name == "Olivia"
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FindLastBy does not works expected\ncase: %v\nexpected: %v taken: %v", input, expected, result)
	}
}

func BenchmarkFilterBy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = gotil.FilterBy([]int64{-100, -5, 30, 100}, func(v int64, i int) bool {
			return v > 0
		})
	}
}

func ExampleFilterBy() {
	result := gotil.FilterBy([]int64{-100, -5, 30, 100}, func(v int64, i int) bool {
		return v > 0
	})
	fmt.Println(result)
	// Output: [30 100]
}
