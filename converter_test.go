package gotil_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gotilty/gotil"
)

func TestToFloat64(t *testing.T) {
	// Float64 to Int
	var input_f int64 = 123
	var expected_f float64 = 123.00
	result_f, err_f := gotil.ToFloat64(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %d\nexpected: %f taken: %f error: %e", input_f, expected_f, result_f, err_f)
	}
	// String to Int
	input_s := "123.12"
	var expected_s float64 = 123.12
	result_s, err_s := gotil.ToFloat64(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %f taken: %f error: %e", input_s, expected_s, result_s, err_s)
	}
	// Bool to Int
	input_b := true
	var expected_b float64 = 1
	result_b, err_b := gotil.ToFloat64(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %f taken: %f error: %e", input_s, expected_s, result_s, err_b)
	}
}
func TestToFloat32(t *testing.T) {
	// Float64 to Int
	var input_f int64 = 123
	var expected_f float32 = 123.00
	result_f, err_f := gotil.ToFloat32(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %d\nexpected: %f taken: %f error: %e", input_f, expected_f, result_f, err_f)
	}
	// String to Int
	input_s := "123.12"
	var expected_s float32 = 123.12
	result_s, err_s := gotil.ToFloat32(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %f taken: %f error: %e", input_s, expected_s, result_s, err_s)
	}
	// Bool to Int
	input_b := true
	var expected_b float32 = 1
	result_b, err_b := gotil.ToFloat32(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %f taken: %f error: %e", input_s, expected_s, result_s, err_b)
	}
}
func TestToInt(t *testing.T) {
	// Float64 to Int
	input_f := 123.12
	var expected_f int = 123
	result_f, err_f := gotil.ToInt(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %f\nexpected: %d taken: %d error: %e", input_f, expected_f, result_f, err_f)
	}
	// String to Int
	input_s := "123"
	var expected_s int = 123
	result_s, err_s := gotil.ToInt(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_s)
	}
	// Bool to Int
	input_b := true
	var expected_b int = 1
	result_b, err_b := gotil.ToInt(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_b)
	}
}
func TestToInt8(t *testing.T) {
	// Float64 to Int8
	input_f := 123.12
	var expected_f int8 = 123
	result_f, err_f := gotil.ToInt8(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %f\nexpected: %d taken: %d error: %e", input_f, expected_f, result_f, err_f)
	}
	// String to Int8
	input_s := "123"
	var expected_s int8 = 123
	result_s, err_s := gotil.ToInt8(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_s)
	}
	// Bool to Int8
	input_b := true
	var expected_b int8 = 1
	result_b, err_b := gotil.ToInt8(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_b)
	}
}
func TestToInt16(t *testing.T) {
	// Float64 to Int16
	input_f := 123.12
	var expected_f int16 = 123
	result_f, err_f := gotil.ToInt16(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %f\nexpected: %d taken: %d error: %e", input_f, expected_f, result_f, err_f)
	}
	// String to Int16
	input_s := "123"
	var expected_s int16 = 123
	result_s, err_s := gotil.ToInt16(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_s)
	}
	// Bool to Int16
	input_b := true
	var expected_b int16 = 1
	result_b, err_b := gotil.ToInt16(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_b)
	}
}
func TestToInt32(t *testing.T) {
	// Float64 to Int32
	input_f := 123.12
	var expected_f int32 = 123
	result_f, err_f := gotil.ToInt32(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %f\nexpected: %d taken: %d error: %e", input_f, expected_f, result_f, err_f)
	}
	// String to Int32
	input_s := "123"
	var expected_s int32 = 123
	result_s, err_s := gotil.ToInt32(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_s)
	}
	// Bool to Int32
	input_b := true
	var expected_b int32 = 1
	result_b, err_b := gotil.ToInt32(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_b)
	}
}
func TestToInt64(t *testing.T) {
	// Float64 to Int64
	input_f := 123.12
	var expected_f int64 = 123
	result_f, err_f := gotil.ToInt64(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %f\nexpected: %d taken: %d error: %e", input_f, expected_f, result_f, err_f)
	}
	// String to Int64
	input_s := "123"
	var expected_s int64 = 123
	result_s, err_s := gotil.ToInt64(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_s)
	}
	// Bool to Int64
	input_b := true
	var expected_b int64 = 1
	result_b, err_b := gotil.ToInt64(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_b)
	}
}
func TestToUint64(t *testing.T) {
	// Float64 to Uint64
	input_f := 123.12
	var expected_f uint64 = 123
	result_f, err_f := gotil.ToUint64(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %f\nexpected: %d taken: %d error: %e", input_f, expected_f, result_f, err_f)
	}
	// String to Uint64
	input_s := "123"
	var expected_s uint64 = 123
	result_s, err_s := gotil.ToUint64(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_s)
	}
	// Bool to Uint64
	input_b := true
	var expected_b uint64 = 1
	result_b, err_b := gotil.ToUint64(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToUint64 does not works expected\ncase: %s\nexpected: %d taken: %d error: %e", input_s, expected_s, result_s, err_b)
	}
}

type toStringTestStruct struct {
	name string
	age  int
}

func TestToString(t *testing.T) {
	input_f := 123.12
	var expected_f string = "123.12"
	result_f, err_f := gotil.ToString(input_f)
	if !reflect.DeepEqual(expected_f, result_f) {
		t.Errorf("Convert.ToString does not works expected\ncase: %f\nexpected: %s taken: %s error: %e", input_f, expected_f, result_f, err_f)
	}
	input_s := 123
	var expected_s string = "123"
	result_s, err_s := gotil.ToString(input_s)
	if !reflect.DeepEqual(expected_s, result_s) {
		t.Errorf("Convert.ToString does not works expected\ncase: %d\nexpected: %s taken: %s error: %e", input_s, expected_s, result_s, err_s)
	}
	input_b := true
	var expected_b string = "true"
	result_b, err_b := gotil.ToString(input_b)
	if !reflect.DeepEqual(expected_b, result_b) {
		t.Errorf("Convert.ToString does not works expected\ncase: %t\nexpected: %s taken: %s error: %e", input_b, expected_b, result_b, err_b)
	}
	input_a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var expected_a string = "[1 2 3 4 5 6 7 8 9 10]"
	result_a, err_a := gotil.ToString(input_a)
	if !reflect.DeepEqual(expected_a, result_a) {
		t.Errorf("Convert.ToString does not works expected\ncase: %b\nexpected: %s taken: %s error: %e", input_a, expected_a, result_a, err_a)
	}
	input_m := map[string][]int{"a": {1, 2, 3, 4, 5}, "b": {11, 22, 33, 44, 55}}
	var expected_m string = `map[a:[1 2 3 4 5] b:[11 22 33 44 55]]`
	result_m, err_m := gotil.ToString(input_m)
	if !reflect.DeepEqual(expected_m, result_m) {
		t.Errorf("Convert.ToString does not works expected\ncase: %v\nexpected: %s taken: %s error: %e", input_m, expected_m, result_m, err_m)
	}
	input_struct := toStringTestStruct{name: "gotil", age: 28}
	var expected_struct string = `{gotil 28}`
	result_struct, err_struct := gotil.ToString(input_struct)
	if !reflect.DeepEqual(expected_struct, result_struct) {
		t.Errorf("Convert.ToString does not works expected\ncase: %v\nexpected: %s taken: %s error: %e", input_struct, expected_struct, result_struct, err_struct)
	}
}

func BenchmarkUint64(b *testing.B) {
	for n := 0; n < b.N; n++ {
		gotil.ToUint64("15")
	}
}

func ExampleUint64() {
	result, err := gotil.ToUint64("15")
	if err == nil {
		fmt.Println(result)
		// Output: 15
	}
	return
}
