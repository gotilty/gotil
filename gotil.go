package gotil

import (
	"fmt"
	"reflect"
	"strconv"
)

// If source is string. Fastest and most reliable way to check if string is empty is s != ""
func isAssignedStr[K string](s K) bool {
	return s != ""
}
func isAssignedint[K int | int32 | int64](s K) bool {
	return s != 0
}
func isAssignedFloat[K float32 | float64](s K) bool {
	return s != 0.0
}
func isAssignedBool[K bool](s K) bool {
	return s == true
}
func isAssignedUint[K uint | uint8 | uint16 | uint32 | uint64](s K) bool {
	return s != 0
}

// IsAssigned
func IsAssigned(a interface{}) bool {
	if a == nil {
		return false
	}
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
		return isAssignedBool(val.Bool())
	case reflect.Pointer, reflect.Chan, reflect.Func, reflect.Struct, reflect.UnsafePointer:
		tt := reflect.Zero(val.Type())
		return !reflect.DeepEqual(val, tt)
	case reflect.Array, reflect.Slice:
		return val.Len() > 0
	default:
		return false
	}

}

func isNotAssignedStr[K string](s K) bool {
	return s == ""
}
func isNotAssignedint[K int | int32 | int64](s K) bool {
	return s == 0
}
func isNotAssignedFloat[K float32 | float64](s K) bool {
	return s == 0.0
}
func isNotAssignedBool[K bool](s K) bool {
	return s != true
}
func isNotAssignedUint[K uint | uint8 | uint16 | uint32 | uint64](s K) bool {
	return s == 0
}

// IsAssigned
func IsNotAssigned(a interface{}) bool {
	if a == nil {
		return true
	}
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return isNotAssignedint(val.Int())
	case reflect.Float32, reflect.Float64:
		return isNotAssignedFloat(val.Float())
	case reflect.String:
		return isNotAssignedStr(fmt.Sprintf("%s", a))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return isNotAssignedUint(val.Uint())
	case reflect.Bool:
		return isNotAssignedBool(val.Bool())
	case reflect.Pointer, reflect.Chan, reflect.Func, reflect.Struct, reflect.UnsafePointer:
		tt := reflect.Zero(val.Type())
		return reflect.DeepEqual(val, tt)
	case reflect.Array, reflect.Slice:
		return val.Len() == 0
	default:
		return false
	}

}

func strToStr[K string](s K) string {
	return fmt.Sprintf("%s", s)
}
func intToStr[K int | int32 | int64](s K) string {
	return strconv.Itoa(int(s))
}
func floatToStr[K float32 | float64](s K) string {
	return fmt.Sprintf("%f", s)
}
func uintToStr[K uint | uint8 | uint16 | uint32 | uint64](s K) string {
	return strconv.FormatUint(uint64(s), 10)
}
func boolToStr[K bool](s K) string {
	if s == true {
		return "true"
	} else {
		return "false"
	}
}

// ToString
func ToString(a interface{}) string {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intToStr(val.Int())
	case reflect.Float32, reflect.Float64:
		return floatToStr(val.Float())
	case reflect.String:
		return strToStr(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintToStr(val.Uint())
	case reflect.Bool:
		return boolToStr(val.Bool())
	case reflect.Pointer, reflect.Chan, reflect.Func, reflect.Struct, reflect.UnsafePointer:
		return val.String()
	case reflect.Array, reflect.Slice:
		var ret string = ""
		var _for int = val.Len()
		for i := 0; i < _for; i++ {
			val := val.Index(i).Interface()
			if i == _for-1 {
				ret += fmt.Sprintf("%s", val)
			} else {
				ret += fmt.Sprintf("%s", val) + ", "
			}

		}
		return ret
	default:
		return "aa"
	}
}
