package converter

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToUint64 returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint64(a interface{}) (uint64, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return uint64(val.Int()), nil
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToUint64(val.Float())
	case reflect.String:
		return fromStringToUint64(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToUint64(val.Uint())
	case reflect.Bool:
		return fromBoolToUint64(val.Bool())
	default:
		if a == nil {
			return 0, nil
		}
		return 0, errs.Int64Error(val.Kind().String())
	}
}

func fromFloat64ToUint64(s float64) (uint64, error) {
	var maxUint uint64 = math.MaxUint64
	if s > float64(maxUint) {
		return 0, errs.CustomError("the entered number cannot be greater than max uint64.")
	}
	return uint64(s), nil
}

func fromStringToUint64(s string) (uint64, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return uint64(d), nil
	}
	return 0, err
}

func fromUint64ToUint64(s uint64) (uint64, error) {
	var maxInt uint64 = math.MaxUint64
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max uint64.")
	}
	return uint64(s), nil
}

func fromBoolToUint64(s bool) (uint64, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
