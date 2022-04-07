package operator_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil/operator"
)

func TestIsArray(t *testing.T) {
	testData := getTestData()
	key := "is_array"
	test := testData[key]
	a := test.output
	b := operator.IsArray(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsSlice(t *testing.T) {
	testData := getTestData()
	key := "is_slice"
	test := testData[key]
	a := test.output
	b := operator.IsSlice(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsString(t *testing.T) {
	testData := getTestData()
	key := "is_string"
	test := testData[key]
	a := test.output
	b := operator.IsString(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsBool(t *testing.T) {
	testData := getTestData()
	key := "is_bool"
	test := testData[key]
	a := test.output
	b := operator.IsBool(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsInteger(t *testing.T) {
	testData := getTestData()
	key := "is_integer"
	test := testData[key]
	a := test.output
	b := operator.IsInteger(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsFloat(t *testing.T) {
	testData := getTestData()
	key := "is_float"
	test := testData[key]
	a := test.output
	b := operator.IsFloat(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsNumber(t *testing.T) {
	testData := getTestData()
	key := "is_number"
	test := testData[key]
	a := test.output
	b := operator.IsNumber(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsFunc(t *testing.T) {
	testData := getTestData()
	key := "is_func"
	test := testData[key]
	a := test.output
	b := operator.IsFunc(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsStruct(t *testing.T) {
	testData := getTestData()
	key := "is_struct"
	test := testData[key]
	a := test.output
	b := operator.IsStruct(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsPointer(t *testing.T) {
	testData := getTestData()
	key := "is_pointer"
	test := testData[key]
	a := test.output
	b := operator.IsPointer(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsChan(t *testing.T) {
	testData := getTestData()
	key := "is_chan"
	test := testData[key]
	a := test.output
	b := operator.IsChan(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsMap(t *testing.T) {
	testData := getTestData()
	key := "is_map"
	test := testData[key]
	a := test.output
	b := operator.IsMap(test.inputValue)
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", key, a, b)
	}
}

func TestIsEqual(t *testing.T) {
	a := operator.IsEqual("hello gotilty", "hello gotilty")
	b := true
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %v taken: %v", "is_equal", a, b)
	}
}

func getTestData() map[string]struct {
	inputValue interface{}
	output     bool
} {
	testData := map[string]struct {
		inputValue interface{}
		output     bool
	}{
		"is_array": {
			inputValue: [4]int64{-100, -5, 30, 100},
			output:     true,
		},
		"is_slice": {
			inputValue: []int64{-100, -5, 30, 100},
			output:     true,
		},
		"is_string": {
			inputValue: "hello gotilty",
			output:     true,
		},
		"is_bool": {
			inputValue: false,
			output:     true,
		},
		"is_integer": {
			inputValue: int16(2),
			output:     true,
		},
		"is_float": {
			inputValue: float64(1523.2323),
			output:     true,
		},
		"is_number": {
			inputValue: float64(1523.2323),
			output:     true,
		},
		"is_func": {
			inputValue: func() {
				fmt.Println("hello gotilty")
			},
			output: true,
		},
		"is_struct": {
			inputValue: user{
				name: "Martha",
				age:  15,
			},
			output: true,
		},
		"is_pointer": {
			inputValue: &user{
				name: "Martha",
				age:  15,
			},
			output: true,
		},
		"is_chan": {
			inputValue: make(chan int),
			output:     true,
		},
		"is_map": {
			inputValue: map[string]struct {
				test bool
			}{
				"test": {
					test: true,
				},
			},
			output: true,
		},
	}
	return testData
}

type user struct {
	age  int
	name string
}
