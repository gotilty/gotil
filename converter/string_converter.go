package converter

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gotility/gotil/config"
)

type arrayJoinOpt struct {
	// optional arguments
	Value string
}

//ToString returns empty string if the parameter is unsupported type
func ToString(a interface{}) string {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intToStr(val.Int())
	case reflect.Float32, reflect.Float64:
		return floatToStr(val.Float())
	case reflect.String:
		return val.String()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintToStr(val.Uint())
	case reflect.Bool:
		return boolToStr(val.Bool())
	case reflect.Pointer, reflect.Chan, reflect.Func, reflect.Struct, reflect.UnsafePointer:
		return val.String()
	case reflect.Array, reflect.Slice:
		return arrayToStringWithSeperator(val, config.GetDefaultSeperator())
	default:
		val.InterfaceData()
		return ""
	}
}

// ToStringWithSeperator
func Join(a interface{}, seperator string) (string, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		return arrayToStringWithSeperator(val, seperator), nil
	default:
		return "", errors.New(fmt.Sprintf("%s cannot join with seperator", val.Kind().String()))
	}
}
func arrayToStringWithSeperator(val reflect.Value, seperator string) string {
	buffer := ""
	length := val.Len()
	for i := 0; i < length; i++ {
		val := val.Index(i).Interface()
		if i == length-1 {
			buffer += fmt.Sprintf("%s", val)
		} else {
			buffer += fmt.Sprintf("%s", val) + seperator
		}
	}
	return buffer
}

func intToStr(s int64) string {
	return strconv.Itoa(int(s))
}
func floatToStr(s float64) string {
	return fmt.Sprintf("%f", s)
}
func uintToStr(s uint64) string {
	return strconv.FormatUint(uint64(s), 10)
}
func boolToStr(s bool) string {
	if s == true {
		return "true"
	} else {
		return "false"
	}
}
