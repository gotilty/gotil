package converter

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func ToInt32(a interface{}) (int32, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fromInt64ToInt32(val.Int())
	case reflect.Float32, reflect.Float64:
		return fromFloat64ToInt32(val.Float())
	case reflect.String:
		return fromStringToInt32(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return fromUint64ToInt32(val.Uint())
	case reflect.Bool:
		return fromBoolToInt32(val.Bool())
	default:
		return 0, errors.New(fmt.Sprintf("The entered value cannot convert from %s to int32", val.Kind().String()))
	}
}

func fromInt64ToInt32(s int64) (int32, error) {
	var maxInt int64 = math.MaxInt32 + 1
	// maxInt == 2147483648
	if s > maxInt {
		return 0, errors.New("The entered number cannot be greater than max int32.")
	}
	return int32(s), nil
}

func fromFloat64ToInt32(s float64) (int32, error) {
	var maxInt int64 = math.MaxInt32 + 1
	if s > float64(maxInt) {
		return 0, errors.New("The entered number cannot be greater than max int32.")
	}
	return int32(s), nil
}

func fromStringToInt32(s string) (int32, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseInt(s, 10, 32)
	if err == nil {
		return fromInt64ToInt32(d)
	}
	return 0, err
}

func fromUint64ToInt32(s uint64) (int32, error) {
	var maxInt int64 = math.MaxInt32 + 1
	if s > uint64(maxInt) {
		return 0, errors.New("The entered number cannot be greater than max int32.")
	}
	return int32(s), nil
}

func fromBoolToInt32(s bool) (int32, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}
