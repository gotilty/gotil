package gotil

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gotilty/gotil/internal/errs"
)

//ToFloat32 returns float32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToFloat32[T any](val T) (float32, error) {
	var result float32
	convertable, def, _ := getConvertable[T, float32](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToFloat64 returns float64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToFloat64[T any](val T) (float64, error) {
	var result float64
	convertable, def, _ := getConvertable[T, float64](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToInt returns int(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt[T any](val T) (int, error) {
	var result int
	convertable, def, _ := getConvertable[T, int](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToInt8 returns int8(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt8[T any](val T) (int8, error) {
	var result int8
	convertable, def, _ := getConvertable[T, int8](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToInt16 returns int16(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt16[T any](val T) (int16, error) {
	var result int16
	convertable, def, _ := getConvertable[T, int16](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToInt32 returns int32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt32[T any](val T) (int32, error) {
	var result int32
	convertable, def, _ := getConvertable[T, int32](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToInt64 returns int64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt64[T any](val T) (int64, error) {
	var result int64
	convertable, def, _ := getConvertable[T, int64](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToUint returns uint(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint[T any](val T) (uint, error) {
	var result uint
	convertable, def, _ := getConvertable[T, uint](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToUint8 returns uint8(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint8[T any](val T) (uint8, error) {
	var result uint8
	convertable, def, _ := getConvertable[T, uint8](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToUint16 returns uint16(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint16[T any](val T) (uint16, error) {
	var result uint16
	convertable, def, _ := getConvertable[T, uint16](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToUint32 returns uint32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint32[T any](val T) (uint32, error) {
	var result uint32
	convertable, def, _ := getConvertable[T, uint32](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

//ToUint64 returns uint64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint64[T any](val T) (uint64, error) {
	var result uint64
	convertable, def, _ := getConvertable[T, uint64](val)
	result, _ = convertable.convert(val, def)
	return result, nil
}

// PRIVATE

type iconverter[T any, D any] interface {
	convert(a T, d D) (D, error)
}

type stringConverter[D any] struct {
	iconverter[string, D]
}

type intConverter[D any] struct {
	iconverter[int, D]
}

type int8Converter[D any] struct {
	iconverter[int8, D]
}

type int16Converter[D any] struct {
	iconverter[int16, D]
}
type int32Converter[D any] struct {
	iconverter[int32, D]
}

type int64Converter[D any] struct {
	iconverter[int64, D]
}

type uintConverter[D any] struct {
	iconverter[uint, D]
}

type uint8Converter[D any] struct {
	iconverter[uint8, D]
}

type uint16Converter[D any] struct {
	iconverter[uint16, D]
}
type uint32Converter[D any] struct {
	iconverter[uint32, D]
}

type uint64Converter[D any] struct {
	iconverter[uint64, D]
}

type float32Converter[D any] struct {
	iconverter[float32, D]
}

type float64Converter[D any] struct {
	iconverter[float64, D]
}

type boolConverter[D any] struct {
	iconverter[bool, D]
}

func sInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	if d, err := strconv.Atoi(s); err == nil {
		return d, nil
	} else {
		return 0, err
	}
}

func sUInt(s string) (uint64, error) {
	if s == "" {
		return 0, nil
	}
	if d, err := strconv.ParseUint(s, 10, 0); err == nil {
		return d, nil
	} else {
		return 0, err
	}
}
func sFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	if d, err := strconv.ParseFloat(s, 64); err == nil {
		return float64(d), nil
	} else {
		return 0, err
	}
}

func sBool(s string) (bool, error) {
	if s == "" {
		return false, nil
	}
	if d, err := strconv.ParseBool(s); err == nil {
		return d, nil
	} else {
		return false, err
	}
}

func (s stringConverter[D]) convert(a string, d D) (D, error) {
	var res interface{}
	var err error
	res = d
	switch any(d).(type) {
	case int, int8, int16, int32, int64:
		var maxInt64 int64 = math.MaxInt64
		if r, err := sInt(a); err == nil {
			if r > int(maxInt64) {
				return d, errs.CustomError("The entered number cannot be greater than max int64.")
			} else {
				switch any(d).(type) {
				case int:
					res = r
				case int8:
					res = int8(r)
				case int16:
					res = int16(r)
				case int32:
					res = int32(r)
				case int64:
					res = int64(r)
				}
				return res.(D), nil
			}
		} else {
			if err == nil {
				return res.(D), nil
			}
			return res.(D), err
		}
	case uint, uint8, uint16, uint32, uint64:
		var maxUInt64 uint64 = math.MaxUint64
		if r, err := sUInt(a); err == nil {
			if r > maxUInt64 {
				return d, errs.CustomError("The entered number cannot be greater than max float64.")
			} else {
				switch any(d).(type) {
				case uint:
					res = r
				case uint8:
					res = uint8(r)
				case uint16:
					res = uint16(r)
				case uint32:
					res = uint32(r)
				case uint64:
					res = uint64(r)
				}
				return res.(D), nil
			}
		} else {
			if err == nil {
				return res.(D), nil
			}
			return res.(D), err
		}
	case float32, float64:
		var maxFloat64 float64 = math.MaxFloat64
		if r, err := sFloat(a); err == nil {
			if r > maxFloat64 {
				return d, errs.CustomError("The entered number cannot be greater than max float64.")
			} else {
				switch any(d).(type) {
				case float32:
					res = float32(r)
				case float64:
					res = float64(r)
				}
				return res.(D), nil
			}
		} else {
			if err == nil {
				return res.(D), nil
			}
			return res.(D), err
		}
	case bool:
		if r, err := sBool(a); err == nil {
			res = r
		} else {
			return res.(D), nil
		}
		return res.(D), err
	case string:
		return d, nil
	}
	return d, nil
}

func (s intConverter[D]) convert(a int, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = a
	case int8:
		if a > math.MaxInt8 || a < math.MinInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		if a > math.MaxInt16 || a < math.MinInt16 {
			return d, errs.ErrOutOfRange
		}
		res = int16(a)
	case int32:
		if a > math.MaxInt32 || a < math.MinInt32 {
			return d, errs.ErrOutOfRange
		}
		res = int32(a)
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if a > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		if a > math.MaxUint16 {
			return d, errs.ErrOutOfRange
		}
		res = uint16(a)
	case uint32:
		if a > math.MaxUint32 {
			return d, errs.ErrOutOfRange
		}
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		if float64(a) > math.MaxFloat32 {
			return d, errs.ErrOutOfRange
		}
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil

}

func (s int8Converter[D]) convert(a int8, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		res = a
	case int16:
		res = int16(a)
	case int32:
		res = int32(a)
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		res = uint8(a)
	case uint16:
		res = uint16(a)
	case uint32:
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil

}

func (s int16Converter[D]) convert(a int16, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		if int(a) > math.MaxInt8 || int(a) < math.MinInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		res = a
	case int32:
		res = int32(a)
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if int(a) > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		res = uint16(a)
	case uint32:
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil
}

func (s int32Converter[D]) convert(a int32, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		if int(a) > math.MaxInt8 || int(a) < math.MinInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		if int(a) > math.MaxInt16 || int(a) < math.MinInt16 {
			return d, errs.ErrOutOfRange
		}
		res = int16(a)
	case int32:
		res = a
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if int(a) > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		if int(a) > math.MaxUint16 {
			return d, errs.ErrOutOfRange
		}
		res = uint16(a)
	case uint32:
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil
}

func (s int64Converter[D]) convert(a int64, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		if int(a) > math.MaxInt8 || int(a) < math.MinInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		if int(a) > math.MaxInt16 || int(a) < math.MinInt16 {
			return d, errs.ErrOutOfRange
		}
		res = int16(a)
	case int32:
		if int(a) > math.MaxInt32 || int(a) < math.MinInt32 {
			return d, errs.ErrOutOfRange
		}
		res = a
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if int(a) > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		if int(a) > math.MaxUint16 {
			return d, errs.ErrOutOfRange
		}
		res = uint16(a)
	case uint32:
		if int(a) > math.MaxUint32 {
			return d, errs.ErrOutOfRange
		}
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil
}

func (s uintConverter[D]) convert(a uint, d D) (D, error) {
	var res interface{}
	switch any(d).(type) {
	case int:
		res = a
	case int8:
		if a > math.MaxInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		if a > math.MaxInt16 {
			return d, errs.ErrOutOfRange
		}
		res = int16(a)
	case int32:
		if a > math.MaxInt32 {
			return d, errs.ErrOutOfRange
		}
		res = int32(a)
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if a > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		if a > math.MaxUint16 {
			return d, errs.ErrOutOfRange
		}
		res = uint16(a)
	case uint32:
		if a > math.MaxUint32 {
			return d, errs.ErrOutOfRange
		}
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		if float64(a) > math.MaxFloat32 {
			return d, errs.ErrOutOfRange
		}
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil

}

func (s uint8Converter[D]) convert(a uint8, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		res = int8(a)
	case int16:
		res = int16(a)
	case int32:
		res = int32(a)
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		res = a
	case uint16:
		res = uint16(a)
	case uint32:
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil

}

func (s uint16Converter[D]) convert(a uint16, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		if a > math.MaxInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		res = int16(a)
	case int32:
		res = int32(a)
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if a > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		res = uint16(a)
	case uint32:
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil
}

func (s uint32Converter[D]) convert(a uint32, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		if a > math.MaxInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		if a > math.MaxInt16 {
			return d, errs.ErrOutOfRange
		}
		res = int16(a)
	case int32:
		res = uint32(a)
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if a > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		if a > math.MaxUint16 {
			return d, errs.ErrOutOfRange
		}
		res = uint16(a)
	case uint32:
		res = a
	case uint64:
		res = uint64(a)
	case float32:
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil
}

func (s uint64Converter[D]) convert(a uint64, d D) (D, error) {
	var res interface{}

	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		if a > math.MaxInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		if a > math.MaxInt16 {
			return d, errs.ErrOutOfRange
		}
		res = int16(a)
	case int32:
		if a > math.MaxInt32 {
			return d, errs.ErrOutOfRange
		}
		res = a
	case int64:
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if a > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		if a > math.MaxUint16 {
			return d, errs.ErrOutOfRange
		}
		res = uint16(a)
	case uint32:
		if a > math.MaxUint32 {
			return d, errs.ErrOutOfRange
		}
		res = uint32(a)
	case uint64:
		res = uint64(a)
	case float32:
		res = float32(a)
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil
}

func (s float32Converter[D]) convert(a float32, d D) (D, error) {
	var res interface{}
	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		if a > math.MaxInt8 || a < math.MinInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		if a > math.MaxInt16 || a < math.MinInt16 {
			return d, errs.ErrOutOfRange
		}
		res = int16(a)
	case int32:
		if a > math.MaxInt32 || a < math.MinInt32 {
			return d, errs.ErrOutOfRange
		}
		res = int32(a)
	case int64:
		if a > math.MaxInt64 || a < math.MinInt64 {
			return d, errs.ErrOutOfRange
		}
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if a > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		if a > math.MaxUint16 {
			return d, errs.ErrOutOfRange
		}
		res = uint16(a)
	case uint32:
		if a > math.MaxUint32 {
			return d, errs.ErrOutOfRange
		}
		res = uint32(a)
	case uint64:
		if a > math.MaxUint64 {
			return d, errs.ErrOutOfRange
		}
		res = uint64(a)
	case float32:
		res = a
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil

}

func (s float64Converter[D]) convert(a float64, d D) (D, error) {
	var res interface{}
	switch any(d).(type) {
	case int:
		res = int(a)
	case int8:
		if a > math.MaxInt8 || a < math.MinInt8 {
			return d, errs.ErrOutOfRange
		}
		res = int8(a)
	case int16:
		if a > math.MaxInt16 || a < math.MinInt16 {
			return d, errs.ErrOutOfRange
		}
		res = int16(a)
	case int32:
		if a > math.MaxInt32 || a < math.MinInt32 {
			return d, errs.ErrOutOfRange
		}
		res = int32(a)
	case int64:
		if a > math.MaxInt64 || a < math.MinInt64 {
			return d, errs.ErrOutOfRange
		}
		res = int64(a)
	case uint:
		res = uint(a)
	case uint8:
		if a > math.MaxUint8 {
			return d, errs.ErrOutOfRange
		}
		res = uint8(a)
	case uint16:
		if a > math.MaxUint16 {
			return d, errs.ErrOutOfRange
		}
		res = uint16(a)
	case uint32:
		if a > math.MaxUint32 {
			return d, errs.ErrOutOfRange
		}
		res = uint32(a)
	case uint64:
		if a > math.MaxUint64 {
			return d, errs.ErrOutOfRange
		}
		res = uint64(a)
	case float32:
		if a > math.MaxFloat32 {
			return d, errs.ErrOutOfRange
		}
		res = a
	case float64:
		res = float64(a)
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a == 1
	default:
		return d, nil
	}
	return res.(D), nil

}

func (s boolConverter[D]) convert(a bool, d D) (D, error) {
	var res interface{}
	switch any(d).(type) {
	case int:
		if a {
			res = 1
		} else {
			res = 0
		}
	case int8:
		if a {
			res = int8(1)
		} else {
			res = int8(0)
		}
	case int16:
		if a {
			res = int16(1)
		} else {
			res = int16(0)
		}
	case int32:
		if a {
			res = int32(1)
		} else {
			res = int32(0)
		}
	case int64:
		if a {
			res = int64(1)
		} else {
			res = int64(0)
		}
	case uint:
		if a {
			res = uint(1)
		} else {
			res = uint(0)
		}
	case uint8:
		if a {
			res = uint8(1)
		} else {
			res = uint8(0)
		}
	case uint16:
		if a {
			res = uint16(1)
		} else {
			res = uint16(0)
		}
	case uint32:
		if a {
			res = uint32(1)
		} else {
			res = uint32(0)
		}
	case uint64:
		if a {
			res = uint64(1)
		} else {
			res = uint64(0)
		}
	case float32:
		if a {
			res = float32(1)
		} else {
			res = float32(0)
		}
	case float64:
		if a {
			res = float64(1)
		} else {
			res = float64(0)
		}
	case string:
		res = fmt.Sprint(a)
	case bool:
		res = a
	default:
		return d, nil
	}
	return res.(D), nil

}

func getConvertable[T, D any](t T) (iconverter[T, D], D, error) {
	var def D
	var convertable interface{}
	switch any(t).(type) {
	case string:
		convertable = stringConverter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case float32:
		convertable = float32Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case float64:
		convertable = float64Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case int:
		convertable = intConverter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case int8:
		convertable = int8Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case int16:
		convertable = int16Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case int32:
		convertable = int32Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case int64:
		convertable = int64Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case uint:
		convertable = uint8Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case uint8:
		convertable = uint8Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case uint16:
		convertable = uint16Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case uint32:
		convertable = uint32Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case uint64:
		convertable = uint64Converter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	case bool:
		convertable = boolConverter[D]{}
		return convertable.(iconverter[T, D]), def, nil
	default:
		return nil, def, errs.ErrUnsupportedType
	}

}
