package converter

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToUint16 returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint16(a interface{}) (uint16, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToUint16(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToUint16(val.Float())
	case reflect.String:
		return fromStringToUint16(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToUint16(val.Uint())
	case reflect.Bool:
		return fromBoolToUint16(val.Bool())
	default:
		if a == nil {
			return 0, nil
		}
		return 0, errs.Int64Error(val.Kind().String())
	}
}

func fromInt64ToUint16(s int64) (uint16, error) {
	var maxInt int64 = math.MaxUint16 + 1
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max uint16.")
	}
	return uint16(s), nil
}

func fromFloat64ToUint16(s float64) (uint16, error) {
	var maxUint uint64 = math.MaxUint16 + 1
	if s > float64(maxUint) {
		return 0, errs.CustomError("the entered number cannot be greater than max uint16.")
	}
	return uint16(s), nil
}

func fromStringToUint16(s string) (uint16, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseUint(s, 10, 16)
	if err == nil {
		return uint16(d), nil
	}
	return 0, err
}

func fromUint64ToUint16(s uint64) (uint16, error) {
	var maxInt uint64 = math.MaxUint16 + 1
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max uint16.")
	}
	return uint16(s), nil
}

func fromBoolToUint16(s bool) (uint16, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
