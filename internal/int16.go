package internal

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToInt16 returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt16(a interface{}) (int16, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToInt16(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToInt16(val.Float())
	case reflect.String:
		return fromStringToInt16(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToInt16(val.Uint())
	case reflect.Bool:
		return fromBoolToInt16(val.Bool())
	default:
		return 0, errs.Int16Error(val.Kind().String())
	}
}

func fromInt64ToInt16(s int64) (int16, error) {
	var maxInt int64 = math.MaxInt16 + 1
	// maxInt == 2147483648
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max int16.")
	}
	return int16(s), nil
}

func fromFloat64ToInt16(s float64) (int16, error) {
	var maxInt int64 = math.MaxInt16 + 1
	if s > float64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int16.")
	}
	return int16(s), nil
}

func fromStringToInt16(s string) (int16, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseInt(s, 10, 16)
	if err == nil {
		return fromInt64ToInt16(d)
	}
	return 0, err
}

func fromUint64ToInt16(s uint64) (int16, error) {
	var maxInt int64 = math.MaxInt16 + 1
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int16.")
	}
	return int16(s), nil
}

func fromBoolToInt16(s bool) (int16, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
