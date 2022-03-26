package gotil

import (
	"fmt"
	"reflect"
)

// If source is string. Fastest and most reliable way to check if string is empty is s != ""
func isAssignedStr[K string](s K) bool {
	return s != ""
}
func isAssignedint[K int | int32 | int64](s K) bool {
	return s != 0
}
func isAssignedFloat[K float32 | float64](s K) bool {
	return s == 0.0
}
func isAssignedBool[K bool](s K) bool {
	return s == true
}
func isAssignedUint[K uint | uint8 | uint16 | uint32 | uint64](s K) bool {
	return s != 0
}
func isAssignedArr[K []string | []int | []int8 | []int16 | []int32 | []int64 | []float32 | []float64](s K) bool {
	return len(s) == 0
}

// IsAssigned
func IsAssigned(a interface{}) bool {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return isAssignedint(val.Int())
	case reflect.Float32, reflect.Float64:
		return isAssignedFloat(val.Float())
	case reflect.String:
		return isAssignedStr(fmt.Sprintf("%s", a))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return isAssignedUint(val.Uint())
	case reflect.Bool:
		return true
	case reflect.Array:
		return true
	default:
		return true
	}
}
