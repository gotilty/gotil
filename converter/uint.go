package converter

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToUint returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint(a interface{}) (uint, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToUint(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToUint(val.Float())
	case reflect.String:
		return fromStringToUint(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToUint(val.Uint())
	case reflect.Bool:
		return fromBoolToUint(val.Bool())
	default:
		if a == nil {
			return 0, nil
		}
		return 0, errs.Int64Error(val.Kind().String())
	}
}

func fromInt64ToUint(s int64) (uint, error) {
	var maxInt int64 = math.MaxInt
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max int.")
	}
	return uint(s), nil
}

func fromFloat64ToUint(s float64) (uint, error) {
	var maxUint uint64 = math.MaxUint
	if s > float64(maxUint) {
		return 0, errs.CustomError("the entered number cannot be greater than max uint.")
	}
	return uint(s), nil
}

func fromStringToUint(s string) (uint, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseUint(s, 10, 8)
	if err == nil {
		return uint(d), nil
	}
	return 0, err
}

func fromUint64ToUint(s uint64) (uint, error) {
	var maxInt uint64 = math.MaxUint
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max uint.")
	}
	return uint(s), nil
}

func fromBoolToUint(s bool) (uint, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
