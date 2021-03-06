package gotil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

type user struct {
	name     string
	age      int
	location location
}

type location struct {
	city string
}

func TestSort(t *testing.T) {

	a := "10"
	r, _ := gotil.ToInt8(a)
	fmt.Println(r)
	// input := []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55}
	// expected := []int64{-100, -5, 5, 11, 30, 33, 55, 100, 1000}
	// result := gotil.Sort(input)
	// if !reflect.DeepEqual(expected, result) {
	// 	t.Errorf("FindLastBy does not works expected\ncase: %d\nexpected: %d taken: %d", input, expected, result)
	// }
}

func ExampleSort() {
	data := []int64{100, 30, -100, -5}
	// Input: [100 30 -100 -5]
	newData := gotil.Sort(data)
	fmt.Println(newData)
	// Output: [-100 -5 30 100]
}

func TestSortBy(t *testing.T) {
	// Input: [{Micheal 27 {New York}} {Joe 30 {Detroit}} {Olivia 42 {New York}} {Kevin 10 {Boston}}]
	input := []user{
		{
			name: "Micheal",
			age:  27,
			location: location{
				city: "New York",
			},
		},
		{
			name: "Joe",
			age:  30,
			location: location{
				city: "Detroit",
			},
		},
		{
			name: "Olivia",
			age:  42,
			location: location{
				city: "New York",
			},
		},
		{
			name: "Kevin",
			age:  10,
			location: location{
				city: "Boston",
			},
		},
	}
	expected := []user{
		{
			name: "Kevin",
			age:  10,
			location: location{
				city: "Boston",
			},
		},

		{
			name: "Joe",
			age:  30,
			location: location{
				city: "Detroit",
			},
		},
		{
			name: "Olivia",
			age:  42,
			location: location{
				city: "New York",
			},
		},
		{
			name: "Micheal",
			age:  27,
			location: location{
				city: "New York",
			},
		},
	}
	result := gotil.SortBy(input, "location.city")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FindLastBy does not works expected\ncase: %v\nexpected: %v taken: %v", input, expected, result)
	}
}

func TestSortDesc(t *testing.T) {
	input := []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55}
	expected := []int64{1000, 100, 55, 33, 30, 11, 5, -5, -100}
	result := gotil.SortDesc(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FindLastBy does not works expected\ncase: %d\nexpected: %d taken: %d", input, expected, result)
	}
}

func ExampleSortDesc() {

	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData := gotil.SortDesc(data)
	fmt.Println(newData)
	// Output: [100 30 -5 -100]
}

func TestSortDescBy(t *testing.T) {
	input := []user{
		{
			name: "Micheal",
			age:  27,
			location: location{
				city: "New York",
			},
		},
		{
			name: "Joe",
			age:  30,
			location: location{
				city: "Detroit",
			},
		},
		{
			name: "Olivia",
			age:  42,
			location: location{
				city: "New York",
			},
		},
		{
			name: "Kevin",
			age:  10,
			location: location{
				city: "Boston",
			},
		},
	}
	expected := []user{
		{
			name: "Olivia",
			age:  42,
			location: location{
				city: "New York",
			},
		},
		{
			name: "Joe",
			age:  30,
			location: location{
				city: "Detroit",
			},
		},
		{
			name: "Micheal",
			age:  27,
			location: location{
				city: "New York",
			},
		},
		{
			name: "Kevin",
			age:  10,
			location: location{
				city: "Boston",
			},
		},
	}
	result := gotil.SortDescBy(input, "age")
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("FindLastBy does not works expected\ncase: %v\nexpected: %v taken: %v", input, expected, result)
	}
}
