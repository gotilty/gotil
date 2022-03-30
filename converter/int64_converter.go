package converter

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func ToInt64(val reflect.Value) (int64, error) {
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToInt64(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToInt64(val.Float())
	case reflect.String:
		return fromStringToInt64(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToInt64(val.Uint())
	case reflect.Bool:
		return fromBoolToInt64(val.Bool())
	default:
		return 0, errors.New(fmt.Sprintf("The entered value cannot convert from %s to int32", val.Kind().String()))
	}
}

func fromInt64ToInt64(s int64) (int64, error) {
	var maxInt int64 = math.MaxInt64
	// maxInt == 2147483648
	if s > maxInt {
		return 0, errors.New("The entered number cannot be greater than max int64.")
	}
	return int64(s), nil
}

func fromFloat64ToInt64(s float64) (int64, error) {
	var maxInt int64 = math.MaxInt64
	if s > float64(maxInt) {
		return 0, errors.New("The entered number cannot be greater than max int64.")
	}
	return int64(s), nil
}

func fromStringToInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return fromInt64ToInt64(d)
	}
	return 0, err
}

func fromUint64ToInt64(s uint64) (int64, error) {
	var maxInt int64 = math.MaxInt32 + 1
	if s > uint64(maxInt) {
		return 0, errors.New("The entered number cannot be greater than max int64.")
	}
	return int64(s), nil
}
func fromBoolToInt64(s bool) (int64, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
