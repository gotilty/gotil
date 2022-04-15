package gotil_test

import (
	"fmt"
	"testing"

	"github.com/gotilty/gotil"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	a := assert.New(t)
	input := []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55}
	expected := []int64{-100, -5, 5, 11, 30, 33, 55, 100, 1000}
	result := gotil.Sort(input)
	a.Equal(expected, result)
}

func ExampleSort() {
	data := []int64{100, 30, -100, -5}
	// Input: [100 30 -100 -5]
	newData := gotil.Sort(data)
	fmt.Println(newData)
	// Output: [-100 -5 30 100]
}

func TestSortDesc(t *testing.T) {
	a := assert.New(t)
	input := []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55}
	expected := []int64{1000, 100, 55, 33, 30, 11, 5, -5, -100}
	result := gotil.SortDesc(input)
	a.Equal(expected, result)
}

func ExampleSortDesc() {

	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData := gotil.SortDesc(data)
	fmt.Println(newData)
	// Output: [100 30 -5 -100]
}
