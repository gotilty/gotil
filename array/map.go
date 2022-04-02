package array

import (
	"reflect"

	. "github.com/gotilty/gotil/config"
	"github.com/gotilty/gotil/converter"
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
func Map(a interface{}, f func(i interface{}, a int) interface{}) (interface{}, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		typ := reflect.TypeOf(a).Elem().Kind()
		if typ == reflect.Int || typ == reflect.Int8 || typ == reflect.Int16 || typ == reflect.Int32 || typ == reflect.Int64 {
			for i := 0; i < val.Len(); i++ {
				arr_i := val.Index(i)
				result, _ := converter.ToInt64(f(arr_i.Int(), i))
				Types.Int64_slice = append(Types.Int64_slice, result)
			}
			return Types.Int64_slice, nil
		} else if typ == reflect.Uint || typ == reflect.Uint8 || typ == reflect.Uint16 || typ == reflect.Uint32 || typ == reflect.Uint64 {
			for i := 0; i < val.Len(); i++ {
				arr_i := val.Index(i)
				result, _ := converter.ToUint64(f(arr_i.Uint(), i))
				Types.Uint64_slice = append(Types.Uint64_slice, result)
			}
			return Types.Uint64_slice, nil
		} else if typ == reflect.String {
			for i := 0; i < val.Len(); i++ {
				arr_i := val.Index(i)
				result, _ := converter.ToString(f(arr_i.String(), i))
				Types.String_slice = append(Types.String_slice, result)
			}
			return Types.String_slice, nil
		}

	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}
