package gotil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	a := assert.New(t)
	expected := 65
	result := Reduce([]int{15, 50}, func(accumulator, v, i int) int {
		return accumulator + v
	}, 0)
	if a.NotEqual(expected, result) {
		t.Errorf("TestReducedoes not works expected\nexpected: %d taken: %d", expected, result)
	}
}

func BenchmarkReduce(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Reduce([]int{15, 50}, func(accumulator, v, i int) int {
			return accumulator + v
		}, 0)
	}
}

func ExampleReduce() {
	result := Reduce([]int{5, 10}, func(accumulator, v, i int) int {
		return accumulator + v
	}, 0)
	fmt.Println(result)
	// Output: 15
}
