package gotil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

func TestEveryBy(t *testing.T) {
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
	expected := true
	result := gotil.EveryBy(input, func(v user) bool {
		return v.age < 50
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("EveryBy does not works expected\ncase: %v\nexpected: %v taken: %v", input, expected, result)
	}
}

func BenchmarkEveryBy(b *testing.B) {
	for n := 0; n < b.N; n++ {
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
		gotil.EveryBy(input, func(v user) bool {
			return v.name == "Olivia"
		})
	}
}

func ExampleEveryBy() {
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
	result := gotil.EveryBy(input, func(v user) bool {
		return v.age < 30
	})
	fmt.Println(result)
	// Output: false
}

func TestContains(t *testing.T) {
	input := []int64{-100, 20, 100}
	expected := false
	result := gotil.Contains(input, 15)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Contains does not works expected\ncase: %v\nexpected: %v taken: %v", input, expected, result)
	}
}

func BenchmarkContains(b *testing.B) {
	for n := 0; n < b.N; n++ {
		input := []int64{-100, 20, 100}
		gotil.Contains(input, 20)
	}
}

func ExampleContains() {
	result := gotil.Contains([]int{5, 10, 20, 30, 50}, 50)
	fmt.Println(result)
	// Output: true
}

func TestContainsBy(t *testing.T) {
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
	expected := true
	result := gotil.ContainsBy(input, func(v user) bool {
		return v.age == 30
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("EveryBy does not works expected\ncase: %v\nexpected: %v taken: %v", input, expected, result)
	}
}

func BenchmarkContainsBy(b *testing.B) {
	for n := 0; n < b.N; n++ {
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
		gotil.ContainsBy(input, func(v user) bool {
			return v.name == "Olivia"
		})
	}
}

func ExampleContainsBy() {
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
	result := gotil.ContainsBy(input, func(v user) bool {
		return v.name == "Mike"
	})
	fmt.Println(result)
	// Output: false
}
