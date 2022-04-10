package internal

import (
	"fmt"
	"reflect"
)

// IsNotAssigned checks the given parameter has not value also if it's slice or array that checks has not any element.
// It works same IsEmpty()
//
// Example:
//
// result := IsNotAssigned("test data")
//
// Output: false
func IsNotAssigned(a interface{}) bool {
	if a == nil {
		return true
	}
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return (val.Int()) == 0
	case reflect.Float32, reflect.Float64:
		return (val.Float()) == 0.0
	case reflect.String:
		return (fmt.Sprintf("%s", a)) == ""
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return (val.Uint()) == 0
	case reflect.Bool:
		return !(val.Bool())
	case reflect.Chan, reflect.Func, reflect.Struct, reflect.UnsafePointer:
		tt := reflect.Zero(val.Type())
		return reflect.DeepEqual(val, tt)
	case reflect.Array, reflect.Slice:
		return val.Len() == 0
	default:
		return false
	}

}
