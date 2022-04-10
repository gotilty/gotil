package internal

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

// Size returns a length of given parameter
// 		data := []int64{100, 30, -100, -5}
// 		result, _ := Size(data)
// 		result2, _ := Size("gotil")
// 		// Output: 4
// 		// Output2: 5
func Size(val interface{}) (int, error) {
	value := reflect.ValueOf(val)
	switch value.Kind() {
	case reflect.Map, reflect.Struct, reflect.Array, reflect.Slice:
		return value.Len(), nil
	case reflect.String:
		val_str, err := ToString(val)
		if err == nil {
			return len(val_str), nil
		}
		return 0, errs.CustomError(value.Kind().String())
	default:
		return 0, errs.CustomError(value.Kind().String())
	}
}
