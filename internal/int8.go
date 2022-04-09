package internal

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToInt8 returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt8(a interface{}) (int8, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToInt8(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToInt8(val.Float())
	case reflect.String:
		return fromStringToInt8(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToInt8(val.Uint())
	case reflect.Bool:
		return fromBoolToInt8(val.Bool())
	default:
		return 0, errs.Int16Error(val.Kind().String())
	}
}

func fromInt64ToInt8(s int64) (int8, error) {
	var maxInt int64 = math.MaxInt16 + 1
	// maxInt == 2147483648
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max int8.")
	}
	return int8(s), nil
}

func fromFloat64ToInt8(s float64) (int8, error) {
	var maxInt int64 = math.MaxInt8 + 1
	if s > float64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int8.")
	}
	return int8(s), nil
}

func fromStringToInt8(s string) (int8, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseInt(s, 10, 8)
	if err == nil {
		return fromInt64ToInt8(d)
	}
	return 0, err
}

func fromUint64ToInt8(s uint64) (int8, error) {
	var maxInt int64 = math.MaxInt8 + 1
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int8.")
	}
	return int8(s), nil
}

func fromBoolToInt8(s bool) (int8, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
