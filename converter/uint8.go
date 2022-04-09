package converter

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToUint8 returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint8(a interface{}) (uint8, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToUint8(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToUint8(val.Float())
	case reflect.String:
		return fromStringToUint8(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToUint8(val.Uint())
	case reflect.Bool:
		return fromBoolToUint8(val.Bool())
	default:
		if a == nil {
			return 0, nil
		}
		return 0, errs.Int64Error(val.Kind().String())
	}
}

func fromInt64ToUint8(s int64) (uint8, error) {
	var maxInt int64 = math.MaxInt8 + 1
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max uint8.")
	}
	return uint8(s), nil
}

func fromFloat64ToUint8(s float64) (uint8, error) {
	var maxUint uint64 = math.MaxUint8 + 1
	if s > float64(maxUint) {
		return 0, errs.CustomError("the entered number cannot be greater than max uint8.")
	}
	return uint8(s), nil
}

func fromStringToUint8(s string) (uint8, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseUint(s, 10, 8)
	if err == nil {
		return uint8(d), nil
	}
	return 0, err
}

func fromUint64ToUint8(s uint64) (uint8, error) {
	var maxInt uint64 = math.MaxUint16 + 1
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max uint8.")
	}
	return uint8(s), nil
}

func fromBoolToUint8(s bool) (uint8, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
