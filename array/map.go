package array

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

// Example Usage
// ----------------------------------------------------------------
// func main() {
// 	sample_array1 := []uint64{9, 12, 15, 34525423452435234, 4, 4532584375937459345, 3, 1, 345345342342423523, 21, 12}
// 	sample_array2 := []string{"Gotilty", "Gotil"}
// 	sample_res1 := array.Map(sample_array1, square).([]uint64)
// 	sample_res2 := array.Map(sample_array2, addChar).([]string)
// 	fmt.Println(sample_res1)
// 	for _, a := range sample_res2 {
// 		fmt.Println(a)
// 	}
// }

// func square(a interface{}, i int) interface{} {
// 	b, _ := converter.ToUint64(a)
// 	return b * 2
// }

// func addChar(a interface{}, i int) interface{} {
// 	b, _ := converter.ToString(a)
// 	if i == 0 {
// 		return b + "->"
// 	}
// 	return b + " is the best utility tool for Golang!"
// }

// Creates an array of values by running each element in the given array thru iteratee.
// The value to be iterated should be given as the first parameter.
// The second parameter will be a function that will take the parameter and return value type interface{}.
// Example: func square(a interface{}, i int()) interface{} {
// 	b, _ := converter.ToInt64(a)
// 	return (b * b)
// }
//
func Map(a interface{}, f func(a interface{}, i int) interface{}) (interface{}, error) {
	val := reflect.ValueOf(a)
	// valF := reflect.ValueOf(f)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		var slice reflect.Value
		var kind reflect.Kind
		assign := false
		for i := 0; i < val.Len(); i++ {
			value := val.Index(i)
			result := f(value.Interface(), i)
			if !assign {
				t2 := reflect.ValueOf(result).Type()
				kind = t2.Kind()
				slice = reflect.Zero(reflect.SliceOf(t2))
				assign = true
			}
			slice = reflect.Append(slice, reflect.ValueOf(result))
		}
		switch kind {
		case reflect.Int:
			return slice.Interface().([]int), nil
		case reflect.Int16:
			return slice.Interface().([]int16), nil
		case reflect.Int32:
			return slice.Interface().([]int32), nil
		case reflect.Int64:
			return slice.Interface().([]int64), nil
		case reflect.String:
			return slice.Interface().([]string), nil
		case reflect.Uint:
			return slice.Interface().([]uint), nil
		case reflect.Uint16:
			return slice.Interface().([]uint16), nil
		case reflect.Uint32:
			return slice.Interface().([]uint32), nil
		case reflect.Uint64:
			return slice.Interface().([]uint64), nil
		case reflect.Float32:
			return slice.Interface().([]float32), nil
		case reflect.Float64:
			return slice.Interface().([]float64), nil
		}
		return slice.Slice, nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}
