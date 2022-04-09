package converter

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToInt returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt(a interface{}) (int, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToInt(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToInt(val.Float())
	case reflect.String:
		return fromStringToInt(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToInt(val.Uint())
	case reflect.Bool:
		return fromBoolToInt(val.Bool())
	default:
		return 0, errs.Int16Error(val.Kind().String())
	}
}

func fromInt64ToInt(s int64) (int, error) {
	var maxInt int64 = math.MaxInt
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max int.")
	}
	return int(s), nil
}

func fromFloat64ToInt(s float64) (int, error) {
	var maxInt int64 = math.MaxInt
	if s > float64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int.")
	}
	return int(s), nil
}

func fromStringToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return fromInt64ToInt(d)
	}
	return 0, err
}

func fromUint64ToInt(s uint64) (int, error) {
	var maxInt int64 = math.MaxInt
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int.")
	}
	return int(s), nil
}

func fromBoolToInt(s bool) (int, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
