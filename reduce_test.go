package gotil_test

import (
	"fmt"
	"testing"

	"github.com/gotilty/gotil"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	a := assert.New(t)
	input := []int{15, 50}
	expected := 65
	result := gotil.Reduce(input, func(accumulator, v, i int) int {
		return accumulator + v
	}, 0)
	a.Equal(expected, result)
}

func BenchmarkReduce(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gotil.Reduce([]int{15, 50}, func(accumulator, v, i int) int {
			return accumulator + v
		}, 0)
	}
}

func ExampleReduce() {
	result := gotil.Reduce([]int{5, 10}, func(accumulator, v, i int) int {
		return accumulator + v
	}, 0)
	fmt.Println(result)
	// Output: 15
}
