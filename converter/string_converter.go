package converter

import (
	"fmt"
	"reflect"
	"strconv"
)

type arrayJoinOpt struct {
	// optional arguments
	Value string
}

// ToString
func ToString(a ...interface{}) string {
	for index, val := range a {
		switch index {
		case 0: //the first mandatory param
			val, _ = val.(string)
		case 1: // age is optional param
			sep, _ = val.(string)
		}
	}
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
				ret += fmt.Sprintf("%s", val) + sep
			}

		}
		return ret
	default:
		return "aa"
	}
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
