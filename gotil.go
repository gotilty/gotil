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
func isAssignedArr[K uint | uint8 | uint16 | uint32 | uint64](s K) bool {
	return s != 0
}

// IsAssigned
func IsAssigned(a interface{}) bool {
	val := reflect.ValueOf(a)
	switch a.(type) {
	case int, int8, int16, int32, int64:
		return isAssignedint(val.Int())
	case float32, float64:
		return isAssignedFloat(val.Float())
	case string:
		return isAssignedStr(fmt.Sprintf("%s", a))
	case uint, uint8, uint16, uint32, uint64:
		return isAssignedUint(val.Uint())
	case bool:
		return true
	case []string:
		return true
	default:
		return true
	}
}
