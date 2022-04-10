package internal

import (
	"reflect"
)

// IsArray checks the given parameter is an array
func IsArray(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Array)
}

// IsSlice checks the given parameter is a slice
func IsSlice(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Slice)
}

// IsString checks the given parameter is a string
func IsString(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.String)
}

// IsBool checks the given parameter is a bool
func IsBool(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Bool)
}

// IsInteger checks the given parameter is an integer
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, unit64
func IsInteger(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKindMultiple(
		val.Kind(),
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
	)
}

// IsFloat checks the given parameter is an float
// float32, float64
func IsFloat(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKindMultiple(
		val.Kind(),
		reflect.Float32,
		reflect.Float64,
	)
}

// IsNumber checks the given parameter is type of any golang number
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, unit64, float32, float64
func IsNumber(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return IsInteger(a) || checkKindMultiple(
		val.Kind(),
		reflect.Float32,
		reflect.Float64,
	)
}

// IsFunc checks the given parameter is a func
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, unit64, float32, float64
func IsFunc(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Func)
}

// IsStruct checks the given parameter is a struct
func IsStruct(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Struct)
}

// IsPointer checks the given parameter is a pointer
func IsPointer(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKindMultiple(
		val.Kind(),
		reflect.Ptr,
	)
}

// IsChan checks the given parameter is a channel
func IsChan(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Chan)
}

// IsMap checks the given parameter is a map
func IsMap(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Map)
}

// IsByteArray checks the given parameter is a byte array
func IsByteArray(a interface{}) bool {
	if a == nil {
		return false
	}
	return reflect.TypeOf(a) == reflect.TypeOf([]byte(nil))
}

// IsEqual checks the given any 2 parameters are equal
func IsEqual(a interface{}, b interface{}) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil {
		return false
	} else if b == nil {
		return false
	}
	return reflect.DeepEqual(a, b)
}

// IsEmpty checks the given parameter has not value also if it's slice or array that checks has not any element.
// It works same IsNotAssigned()
func IsEmpty(a interface{}, b interface{}) bool {
	return IsNotAssigned(a)
}

// IsNil checks the given parameter has not value also if it's slice or array that checks has not any element.
// It works same IsNotAssigned()
func IsNil(a interface{}, b interface{}) bool {
	return IsNotAssigned(a)
}

func checkKind(k reflect.Kind, t reflect.Kind) bool {
	return k == t
}

func checkKindMultiple(k reflect.Kind, t ...reflect.Kind) bool {
	if r, err := FindBy(t, func(val interface{}, i int) bool {
		if kk, ok := val.(reflect.Kind); ok && kk == k {
			return true
		} else {
			return false
		}
	}); err == nil && r != nil {
		return true
	} else {
		return false
	}
}
