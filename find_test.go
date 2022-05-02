package gotil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

func TestFindBy(t *testing.T) {
	input := []int64{-100, -5, 30, 100}
	expected := int64(30)
	result := gotil.FindBy(input, func(val int64) bool {
		return val == 30
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("F does not works expected\ncase: %d\nexpected: %d taken: %d", input, expected, result)
	}
	input2 := []user{
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
	result2 := gotil.FindBy(input2, func(val user) bool {
		return val.name == "Olivia"
	})
	expected2 := user{
		name: "Olivia",
		age:  42,
	}
	if !reflect.DeepEqual(expected2, result2) {
		t.Errorf("FindBy does not works expected\ncase: %v\nexpected: %v taken: %v", input2, expected2, result2)
	}
}

func BenchmarkFindByIntegerSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := []int64{-100, -5, 30, 100}
		gotil.FindBy(input, func(val int64) bool {
			return val == 30
		})
	}
}

func BenchmarkFindByStructSlice(b *testing.B) {
	input2 := []user{
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
	for n := 0; n < b.N; n++ {
		gotil.FindBy(input2, func(val user) bool {
			return val.name == "Olivia"
		})
	}
}

func ExampleFindBy() {
	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData := gotil.FindBy(data, func(val int64) bool {
		return val == 30
	})
	fmt.Println(newData)
	// Output: 30
}

func TestFindLastBy(t *testing.T) {
	input := []int64{10, 15, 90, 100, 30, 50}
	expected := int64(100)
	result := gotil.FindLastBy(input, func(val int64) bool {
		return val > 80
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FindLastBy does not works expected\ncase: %d\nexpected: %d taken: %d", input, expected, result)
	}
	input2 := []user{
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
		{
			name: "Mike",
			age:  43,
		},
	}
	result2 := gotil.FindBy(input2, func(val user) bool {
		return val.age == 42
	})
	expected2 := user{
		name: "Olivia",
		age:  42,
	}
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FindLastBy does not works expected\ncase: %v\nexpected: %v taken: %v", input2, expected2, result2)
	}
}

func BenchmarkFindLastByIntegerSlice(b *testing.B) {
	input := []int64{10, 15, 90, 100, 30, 50}
	for n := 0; n < b.N; n++ {
		_ = gotil.FindLastBy(input, func(val int64) bool {
			return val > 80
		})
	}
}

func BenchmarkFindLastByStructSlice(b *testing.B) {
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
		{
			name: "Mike",
			age:  43,
		},
	}
	for n := 0; n < b.N; n++ {
		_ = gotil.FindLastBy(input, func(val user) bool {
			return val.age == 42
		})
	}
}

func ExampleFindLastBy() {
	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData := gotil.FindLastBy(data, func(val int64) bool {
		return val == 30
	})
	fmt.Println(newData)
	// Output: 30
}
