package internal

import (
	"reflect"
	"strings"

	"github.com/gotilty/gotil/internal/errs"
)

// It checks whether the given value exists in the given parameter.
// 	data := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}
// 	result, _ := Includes(data, "gotil")
// 	// Output: false
func Includes(a interface{}, val interface{}) (bool, error) {
	state_arr := false
	state_map := false
	value := reflect.ValueOf(a)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < value.Len(); i++ {
			value_i := value.Index(i)
			if reflect.DeepEqual(value_i.Interface(), val) {
				return true, nil
			}
		}
		return state_arr, nil
	case reflect.Map, reflect.Struct:
		keys := value.MapKeys()
		FindBy(keys, func(find interface{}, i int) bool {
			slice_in_map := value.MapIndex(find.(reflect.Value))
			if includes_in_slice, err := Includes(slice_in_map.Interface(), val); err == nil && includes_in_slice {
				state_map = includes_in_slice
			}
			return state_map
		})
		return state_map || state_arr, nil

	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		a_str, _ := ToString(a)
		val_str, _ := ToString(val)
		return (strings.Contains(a_str, val_str)), nil
	}
	return false, errs.NewUnsupportedTypeError(value.Kind().String())
}
