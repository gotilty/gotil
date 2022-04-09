package internal

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToFloat64 returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToFloat64(a interface{}) (float64, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToFloat64(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToFloat64(val.Float())
	case reflect.String:
		return fromStringToFloat64(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToFloat64(val.Uint())
	case reflect.Bool:
		return fromBoolToFloat64(val.Bool())
	default:
		if a == nil {
			return 0, nil
		}
		return 0, errs.Int64Error(val.Kind().String())
	}
}

func fromInt64ToFloat64(s int64) (float64, error) {
	var maxInt64 int64 = math.MaxInt64
	if s > maxInt64 {
		return 0, errs.CustomError("The entered number cannot be greater than max float64.")
	}
	return float64(s), nil
}

func fromFloat64ToFloat64(s float64) (float64, error) {
	var maxFloat float64 = math.MaxFloat64
	if s > float64(maxFloat) {
		return 0, errs.CustomError("the entered number cannot be greater than max float64.")
	}
	return float64(s), nil
}

func fromStringToFloat64(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return float64(d), nil
	}
	return 0, err
}

func fromUint64ToFloat64(s uint64) (float64, error) {
	var maxFloat float64 = math.MaxFloat64
	if s > uint64(maxFloat) {
		return 0, errs.CustomError("The entered number cannot be greater than max float64.")
	}
	return float64(s), nil
}
func fromBoolToFloat64(s bool) (float64, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
