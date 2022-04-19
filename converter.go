package gotil

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToFloat32 returns float32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToFloat32[T any](val T) (float32, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToFloat32(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		val_int, err_a := ToInt(val)
		if err_a == nil {
			v, err_b := fromIntToFloat32(val_int)
			if err_b == nil {
				return v, nil
			}
		}
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToFloat32(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToFloat32(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToFloat32(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToFloat64 returns float64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToFloat64[T any](val T) (float64, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToFloat64(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		val_int, err_a := ToInt(val)
		if err_a == nil {
			v, err_b := fromIntToFloat64(val_int)
			if err_b == nil {
				return v, nil
			}
		}

	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToFloat64(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToFloat64(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToFloat64(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToInt returns int(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt[T any](val T) (int, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToInt(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int:
		return int(any(val).(int)), nil
	case int8:
		return int(any(val).(int8)), nil
	case int16:
		return int(any(val).(int16)), nil
	case int32:
		return int(any(val).(int32)), nil
	case int64:
		return int(any(val).(int64)), nil
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToInt(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToInt(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToInt(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToInt8 returns int8(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt8[T any](val T) (int8, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToInt8(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		v, err := fromInt64ToInt8(any(val).(int64))
		if err == nil {
			return v, nil
		}
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToInt8(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToInt8(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToInt8(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToInt16 returns int16(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt16[T any](val T) (int16, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToInt16(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		v, err := fromInt64ToInt16(any(val).(int64))
		if err == nil {
			return v, nil
		}
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToInt16(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToInt16(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToInt16(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToInt32 returns int32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt32[T any](val T) (int32, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToInt32(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		v, err := fromInt64ToInt32(any(val).(int64))
		if err == nil {
			return v, nil
		}
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToInt32(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToInt32(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToInt32(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToInt64 returns int64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt64[T any](val T) (int64, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToInt64(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		return int64(any(val).(int)), nil
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToInt64(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToInt64(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToInt64(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToString returns empty string if the parameter is unsupported type.
//Just works with all primitive types, arrays and slices.
//alternative for slice and array gotil.Join(a, separator)
//
//Example1:
// 		ToString([]int{10,20,30}) returns "10,20,30"
//Example2:
// 	data := []user{
// 		{
// 			name: "Micheal",
// 			age:  27,
// 		},
// 		{
// 			name: "Joe",
// 			age:  30,
// 	 	},
// 	}
// 	ToString(data)
// 	// Output: "{Micheal 27},{Joe 30}"
func ToString[T any](s ...T) (string, error) {
	var buffer bytes.Buffer
	for _, val := range s {
		switch any(val).(type) {
		case string:
			buffer.WriteString(any(val).(string))
		case int, int8, int16, int32, int64:
			v, err := fromIntToStr(any(val).(int))
			if err == nil {
				buffer.WriteString(v)
			}
		case uint, uint8, uint16, uint32, uint64:
			v, err := fromUint64ToStr(any(val).(uint64))
			if err == nil {
				buffer.WriteString(v)
			}
		case float32, float64:
			v, err := fromFloat64ToStr(any(val).(float64))
			if err == nil {
				buffer.WriteString(v)
			}
		case bool:
			buffer.WriteString(fromBoolToStr(any(val).(bool)))
		default:
			if reflect.TypeOf(val).Kind() == reflect.Array ||
				reflect.TypeOf(val).Kind() == reflect.Slice ||
				reflect.TypeOf(val).Kind() == reflect.Map ||
				reflect.TypeOf(val).Kind() == reflect.Struct {
				for _, v := range s {
					buffer.WriteString(fmt.Sprint(v))
				}
			}
		}
		return buffer.String(), nil

	}
	return "", nil

}

//Join just works with array and slice
//
//Example1:
//	Join([]int{10,20,30},"-") returns "10-20-30"
func Join[T any](val []T, seperator string) (string, error) {
	if val == nil {
		return "", errs.NilReferenceTypeError()
	}
	if len(val) == 0 {
		return "", nil
	}
	switch any(val).(type) {
	case []string:
		v, err := arrayToStringWithSeparator(any(val).([]string), seperator)
		if err != nil {
			return v, nil
		}
		return "", err
	case []int, []int8, []int16, []int32, []int64:
		v, err := arrayToStringWithSeparator(any(val).([]int64), seperator)
		if err != nil {
			return v, nil
		}
		return "", err
	case []uint, []uint8, []uint16, []uint32, []uint64:
		v, err := arrayToStringWithSeparator(any(val).([]uint64), seperator)
		if err != nil {
			return v, nil
		}
		return "", err
	case []float32, []float64:
		v, err := arrayToStringWithSeparator(any(val).([]float64), seperator)
		if err != nil {
			return v, nil
		}
		return "", err
	case []bool:
		v, err := arrayToStringWithSeparator(any(val).([]bool), seperator)
		if err != nil {
			return v, nil
		}
		return "", err
	}
	return "", nil
}

//ToUint returns uint(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint[T any](val T) (uint, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToUint(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		v, err := fromInt64ToUint(any(val).(int64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToUint(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToUint(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToUint(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToUint8 returns uint8(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint8[T any](val T) (uint8, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToUint8(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		v, err := fromInt64ToUint8(any(val).(int64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToUint8(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToUint8(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToUint8(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToUint16 returns uint16(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint16[T any](val T) (uint16, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToUint16(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		v, err := fromInt64ToUint16(any(val).(int64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToUint16(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToUint16(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToUint16(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToUint32 returns uint32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint32[T any](val T) (uint32, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToUint32(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		v, err := fromInt64ToUint32(any(val).(int64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToUint32(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToUint32(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToUint32(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

//ToUint64 returns uint64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint64[T any](val T) (uint64, error) {
	switch any(val).(type) {
	case string:
		v, err := fromStringToUint64(any(val).(string))
		if err == nil {
			return v, nil
		}
		return 0, err
	case int, int8, int16, int32, int64:
		return uint64(any(val).(int)), nil
	case uint, uint8, uint16, uint32, uint64:
		v, err := fromUint64ToUint64(any(val).(uint64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case float32, float64:
		v, err := fromFloat64ToUint64(any(val).(float64))
		if err == nil {
			return v, nil
		}
		return 0, err
	case bool:
		v, err := fromBoolToUint64(any(val).(bool))
		if err == nil {
			return v, nil
		}
		return 0, err
	}

	return 0, nil
}

// PRIVATE

func fromIntToFloat32(s int) (float32, error) {
	var maxInt32 float64 = math.MaxInt32
	if s > int(maxInt32) {
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

func fromIntToFloat64(s int) (float64, error) {
	var maxInt64 int = math.MaxInt64
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

func fromInt64ToInt(s int64) (int, error) {
	var maxInt int64 = math.MaxInt
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max int.")
	}
	return int(s), nil
}

func fromFloat64ToInt(s float64) (int, error) {
	var maxInt int64 = math.MaxInt
	if s > float64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int.")
	}
	return int(s), nil
}

func fromStringToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return fromInt64ToInt(d)
	}
	return 0, err
}

func fromUint64ToInt(s uint64) (int, error) {
	var maxInt int64 = math.MaxInt
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int.")
	}
	return int(s), nil
}

func fromBoolToInt(s bool) (int, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}

func fromInt64ToInt8(s int64) (int8, error) {
	var maxInt int64 = math.MaxInt16 + 1
	// maxInt == 2147483648
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max int8.")
	}
	return int8(s), nil
}

func fromFloat64ToInt8(s float64) (int8, error) {
	var maxInt int64 = math.MaxInt8 + 1
	if s > float64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int8.")
	}
	return int8(s), nil
}

func fromStringToInt8(s string) (int8, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseInt(s, 10, 8)
	if err == nil {
		return fromInt64ToInt8(d)
	}
	return 0, err
}

func fromUint64ToInt8(s uint64) (int8, error) {
	var maxInt int64 = math.MaxInt8 + 1
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int8.")
	}
	return int8(s), nil
}

func fromBoolToInt8(s bool) (int8, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}

func fromInt64ToInt16(s int64) (int16, error) {
	var maxInt int64 = math.MaxInt16 + 1
	// maxInt == 2147483648
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max int16.")
	}
	return int16(s), nil
}

func fromFloat64ToInt16(s float64) (int16, error) {
	var maxInt int64 = math.MaxInt16 + 1
	if s > float64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int16.")
	}
	return int16(s), nil
}

func fromStringToInt16(s string) (int16, error) {
	if s == "" {
		return 0, nil
	}
	d, err := strconv.ParseInt(s, 10, 16)
	if err == nil {
		return fromInt64ToInt16(d)
	}
	return 0, err
}

func fromUint64ToInt16(s uint64) (int16, error) {
	var maxInt int64 = math.MaxInt16 + 1
	if s > uint64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int16.")
	}
	return int16(s), nil
}

func fromBoolToInt16(s bool) (int16, error) {
	if s {
		return 1, nil
	} else {
		return 0, nil
	}
}

func fromInt64ToInt32(s int64) (int32, error) {
	var maxInt int64 = math.MaxInt32 + 1
	// maxInt == 2147483648
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max int32.")
	}
	return int32(s), nil
}

func fromFloat64ToInt32(s float64) (int32, error) {
	var maxInt int64 = math.MaxInt32 + 1
	if s > float64(maxInt) {
		return 0, errs.CustomError("The entered number cannot be greater than max int32.")
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
		return 0, errs.CustomError("The entered number cannot be greater than max int32.")
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

func fromFloat64ToInt64(s float64) (int64, error) {
	var maxInt int64 = math.MaxInt64
	if s > float64(maxInt) {
		return 0, errs.CustomError("the entered number cannot be greater than max int64.")
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
		return 0, errs.CustomError("The entered number cannot be greater than max int64.")
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

func fromIntToStr(s int) (string, error) {
	var maxInt int = math.MaxInt
	if s > maxInt {
		return "", errs.CustomError("the entered number cannot be greater than max int.")
	}
	return strconv.Itoa(int(s)), nil
}

func fromFloat64ToStr(s float64) (string, error) {
	var maxFloat float64 = math.MaxFloat64
	if s > maxFloat {
		return "", errs.CustomError("the entered number cannot be greater than max uint.")
	}
	return strconv.FormatFloat(s, 'f', -1, 64), nil
}

func fromUint64ToStr(s uint64) (string, error) {
	var maxUint uint64 = math.MaxUint
	if s > maxUint {
		return "", errs.CustomError("the entered number cannot be greater than max uint.")
	}
	return strconv.FormatUint(uint64(s), 10), nil
}
func fromBoolToStr(s bool) string {
	if s == true {
		return "true"
	} else {
		return "false"
	}
}

func arrayToStringWithSeparator[T any](val []T, separator string) (string, error) {
	buffer := ""
	length := len(val)
	for i := 0; i < length; i++ {
		val := val[i]
		if b, err := ToString(val); err == nil {
			if i == length-1 {
				buffer += b
			} else {
				buffer += b + separator
			}
		} else {
			return "", err
		}
	}
	return buffer, nil
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

func fromFloat64ToUint32(s float64) (uint32, error) {
	var maxUint uint64 = math.MaxUint32 + 1
	if s > float64(maxUint) {
		return 0, errs.CustomError("the entered number cannot be greater than max uint32.")
	}
	return uint32(s), nil
}

func fromInt64ToUint32(s int64) (uint32, error) {
	var maxInt int64 = math.MaxUint32 + 1
	if s > maxInt {
		return 0, errs.CustomError("The entered number cannot be greater than max uint32.")
	}
	return uint32(s), nil
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

func fromFloat64ToUint64(s float64) (uint64, error) {
	var maxUint uint64 = math.MaxUint64
	if s > float64(maxUint) {
		return 0, errs.CustomError("the entered number cannot be greater than max uint64.")
	}
	return uint64(s), nil
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
