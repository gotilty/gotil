package converter

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func ToInt64(a interface{}) (int64, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int64(val.Int()), nil
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToInt64(val.Float())
	case reflect.String:
		return fromStringToInt64(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToInt64(val.Uint())
	case reflect.Bool:
		return fromBoolToInt64(val.Bool())
	default:
		if a == nil {
			return 0, nil
		}
		return 0, errors.New(fmt.Sprintf("The entered value cannot convert from %s to int32", val.Kind().String()))
	}
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
		return int64(d), nil
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
