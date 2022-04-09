package converter

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToUint32 returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint32(a interface{}) (uint32, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToUint32(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToUint32(val.Float())
	case reflect.String:
		return fromStringToUint32(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToUint32(val.Uint())
	case reflect.Bool:
		return fromBoolToUint32(val.Bool())
	default:
		if a == nil {
			return 0, nil
		}
		return 0, errs.Int64Error(val.Kind().String())
	}
}

func fromInt64ToUint32(s int64) (uint32, error) {
	var maxInt int64 = math.MaxUint32 + 1
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max uint32.")
	}
	return uint32(s), nil
}

func fromFloat64ToUint32(s float64) (uint32, error) {
	var maxUint uint64 = math.MaxUint32 + 1
	if s > float64(maxUint) {
		return 0, errs.CustomError("the entered number cannot be greater than max uint32.")
	}
	return uint32(s), nil
}

func fromStringToUint32(s string) (uint32, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseUint(s, 10, 32)
	if err == nil {
		return uint32(d), nil
	}
	return 0, err
}

func fromUint64ToUint32(s uint64) (uint32, error) {
	var maxInt uint64 = math.MaxUint32 + 1
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max uint32.")
	}
	return uint32(s), nil
}

func fromBoolToUint32(s bool) (uint32, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
