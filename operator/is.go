package operator

import (
	"reflect"

	"github.com/gotilty/gotil/collection"
)

func IsArray(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Array)
}

func IsSlice(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Slice)
}

func IsString(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.String)
}

func IsBool(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Bool)
}

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

func IsFunc(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Func)
}

func IsStruct(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Struct)
}

func IsPointer(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKindMultiple(
		val.Kind(),
		reflect.Pointer,
		reflect.UnsafePointer,
	)
}

func IsChan(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Chan)
}

func IsMap(a interface{}) bool {
	if a == nil {
		return false
	}
	val := reflect.ValueOf(a)
	return checkKind(val.Kind(), reflect.Map)
}

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

func checkKind(k reflect.Kind, t reflect.Kind) bool {
	return k == t
}

func checkKindMultiple(k reflect.Kind, t ...reflect.Kind) bool {
	if r, err := collection.FindBy(t, func(val interface{}, i int) bool {
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
