// package collection

// import (
// 	"fmt"
// 	"reflect"

// 	"github.com/gotilty/gotil/internal/errs"
// )

// func GroupBy(a interface{}, f func(val interface{}, i int) interface{}) (interface{}, error) {
// 	val := reflect.ValueOf(a)
// 	switch val.Kind() {
// 	case reflect.Slice, reflect.Array:
// 		var slice reflect.Value
// 		assign := false
// 		for i := 0; i < val.Len(); i++ {
// 			value := val.Index(i)
// 			result := f(value.Interface(), i)
// 			mapType := reflect.MapOf(reflect.TypeOf(result), reflect.TypeOf(a))
// 			fmt.Println(mapType)
// 			mapValue := reflect.MakeMap(mapType)
// 			keys := mapValue.MapKeys()
// 			reflect.Append(keys[1], value)
// 			for i, k := range keys {
// 				c_key := k.Convert(mapValue.Type().Key())
// 				c_value := mapValue.MapIndex(c_key)
// 				fmt.Println(c_value)
// 				if reflect.DeepEqual(c_key, reflect.ValueOf(result)) {
// 					reflect.Append(keys[i], k)
// 				} else {
// 					mapValue.SetMapIndex(reflect.ValueOf(i), c_value)
// 				}
// 			}
// 			fmt.Println(mapValue.Interface())
// 			if !assign {
// 				t2 := value.Type()
// 				slice = reflect.Zero(reflect.SliceOf(t2))
// 				assign = true
// 			}
// 			// slice = reflect.Append(slice, reflect.ValueOf(mapValue))
// 		}

// 		return slice.Interface(), nil
// 	}
// 	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
// }
