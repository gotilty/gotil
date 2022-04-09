package converter

import (
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToFloat32 returns 0 with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToFloat32(a interface{}) (float32, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToFloat32(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToFloat32(val.Float())
	case reflect.String:
		return fromStringToFloat32(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToFloat32(val.Uint())
	case reflect.Bool:
		return fromBoolToFloat32(val.Bool())
	default:
		if a == nil {
			return 0, nil
		}
		return 0, errs.Int64Error(val.Kind().String())
	}
}

func fromInt64ToFloat32(s int64) (float32, error) {
	var maxInt32 float64 = math.MaxInt32
	if s > int64(maxInt32) {
		return 0, errs.CustomError("The entered number cannot be greater than max float32.")
	}
	return float32(s), nil
}

func fromFloat64ToFloat32(s float64) (float32, error) {
	var maxInt float64 = math.MaxFloat32
	if s > float64(maxInt) {
		return 0, errs.CustomError("the entered number cannot be greater than max float32.")
	}
	return float32(s), nil
}

func fromStringToFloat32(s string) (float32, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseFloat(s, 32)
	if err == nil {
		return float32(d), nil
	}
	return 0, err
}

func fromUint64ToFloat32(s uint64) (float32, error) {
	var maxFloat32 float64 = math.MaxFloat32
	if s > uint64(maxFloat32) {
		return 0, errs.CustomError("The entered number cannot be greater than max float32.")
	}
	return float32(s), nil
}
func fromBoolToFloat32(s bool) (float32, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
