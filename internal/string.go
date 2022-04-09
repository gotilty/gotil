package internal

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/config"
	"github.com/gotilty/gotil/internal/errs"
)

type arrayJoinOpt struct {
	// optional arguments
	Value string
}

//ToString returns empty string if the parameter is unsupported type.
//Just works with all primitive types, arrays and slices.
func ToString(a interface{}) (string, error) {
	if a == nil {
		return "", errs.NilReferenceTypeError()
	}
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intToStr(val.Int()), nil
	case reflect.Float32, reflect.Float64:
		return floatToStr(val.Float()), nil
	case reflect.String:
		return val.String(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintToStr(val.Uint()), nil
	case reflect.Bool:
		return boolToStr(val.Bool()), nil
	case reflect.Array, reflect.Slice:
		return arrayToStringWithSeperator(val, config.GetDefaultSeperator())
	case reflect.Map, reflect.Ptr, reflect.Complex128, reflect.Complex64, reflect.Interface, reflect.Struct:
		return fmt.Sprint(val), nil
	default:
		return "", errs.NewUnsupportedTypeError(val.Kind().String())
	}
}

// ToStringWithSeperator
func Join(a interface{}, seperator string) (string, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		return arrayToStringWithSeperator(val, seperator)
	default:
		return "", errors.New(fmt.Sprintf("%s cannot join with seperator", val.Kind().String()))
	}
}
func arrayToStringWithSeperator(val reflect.Value, seperator string) (string, error) {
	buffer := ""
	length := val.Len()
	for i := 0; i < length; i++ {
		val := val.Index(i)
		if b, err := ToString(val.Interface()); err == nil {
			if i == length-1 {
				buffer += b
			} else {
				buffer += b + seperator
			}
		} else {
			return "", err
		}
	}
	return buffer, nil
}

func intToStr(s int64) string {
	return strconv.Itoa(int(s))
}

/*
 1.234560e+02		Scientific notation	%e
 123.456000			Decimal point, no exponent	%f
 123.46				Default width, precision 2	%.2f
 123.46				Width 8, precision 2	%8.2f
 123.456			Exponent as needed, necessary digits only	%g
*/

func floatToStr(s float64) string {
	return strconv.FormatFloat(s, 'f', -1, 64)
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
