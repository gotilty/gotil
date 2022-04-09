package collection_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil/collection"
)

func TestSort(t *testing.T) {
	testData := getSortTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := collection.Sort(test.inputValue)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func ExampleSort() {

	data := []int64{100, 30, -100, -5}
	// Input: [100 30 -100 -5]
	newData, _ := collection.Sort(data)
	fmt.Println(newData)
	// Output: [-100 -5 30 100]
}

func TestSortBy(t *testing.T) {
	testData := getSortByTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := collection.SortBy(test.inputValue, test.property)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func ExampleSortBy() {
	data := []user{
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
	// Input: [{Micheal 27 {New York}} {Joe 30 {Detroit}} {Olivia 42 {New York}} {Kevin 10 {Boston}}]
	newData, _ := collection.SortBy(data, "location.city")
	fmt.Println(newData)
	// Output: [{Kevin 10 {Boston}} {Joe 30 {Detroit}} {Olivia 42 {New York}} {Micheal 27 {New York}}]
}

func TestSortDesc(t *testing.T) {
	testData := getSortDescTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := collection.SortDesc(test.inputValue)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func ExampleSortDesc() {

	data := []int64{-100, -5, 30, 100}
	// Input: [-100 -5 30 100]
	newData, _ := collection.SortDesc(data)
	fmt.Println(newData)
	// Output: [100 30 -5 -100]
}

func TestSortDescBy(t *testing.T) {
	testData := getSortDescByTestData()
	for key, test := range testData {
		a, erra := test.output, test.err
		b, errb := collection.SortDescBy(test.inputValue, test.property)
		if erra == nil {
			if !reflect.DeepEqual(a, b) || errb != nil {
				t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %s", key, a, b, errb.Error())
			}
		}
	}
}

func ExampleSortDescBy() {
	data := []user{
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
	// Input: [{Micheal 27 {New York}} {Joe 30 {Detroit}} {Olivia 42 {New York}} {Kevin 10 {Boston}}]
	newData, _ := collection.SortDescBy(data, "location.city")
	fmt.Println(newData)
	// Output: [{Micheal 27 {New York}} {Olivia 42 {New York}} {Joe 30 {Detroit}} {Kevin 10 {Boston}}]
}

func getSortTestData() map[string]struct {
	inputValue interface{}
	output     interface{}
	property   string
	err        error
} {

	testData := map[string]struct {
		inputValue interface{}
		output     interface{}
		property   string
		err        error
	}{
		"sort_numbers": {
			inputValue: []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55},
			output:     []int64{-100, -5, 5, 11, 30, 33, 55, 100, 1000},
			property:   "",
			err:        nil,
		},
	}
	return testData
}

func getSortDescTestData() map[string]struct {
	inputValue interface{}
	output     interface{}
	property   string
	err        error
} {

	testData := map[string]struct {
		inputValue interface{}
		output     interface{}
		property   string
		err        error
	}{
		"sort_numbers": {
			inputValue: []int64{-100, -5, 30, 100, 5, 11, 1000, 33, 55},
			output:     []int64{1000, 100, 55, 33, 30, 11, 5, -5, -100},
			property:   "",
			err:        nil,
		},
	}
	return testData
}

func getSortByTestData() map[string]struct {
	inputValue interface{}
	output     interface{}
	property   string
	err        error
} {

	testData := map[string]struct {
		inputValue interface{}
		output     interface{}
		property   string
		err        error
	}{
		"sort_users": {
			inputValue: []user{
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
			},
			output: []user{
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
			},
			property: "location.city",
			err:      nil,
		},
	}
	return testData
}

func getSortDescByTestData() map[string]struct {
	inputValue interface{}
	output     interface{}
	property   string
	err        error
} {

	testData := map[string]struct {
		inputValue interface{}
		output     interface{}
		property   string
		err        error
	}{
		"sort_users": {
			inputValue: []user{
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
			},
			output: []user{
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
			},
			property: "age",
			err:      nil,
		},
	}
	return testData
}
