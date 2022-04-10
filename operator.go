package gotil

import "github.com/gotilty/gotil/internal"

// IsAssigned
func IsAssigned(a interface{}) bool {
	return internal.IsAssigned(a)
}

// IsNotAssigned
func IsNotAssigned(a interface{}) bool {
	return internal.IsNotAssigned(a)
}

func IsArray(a interface{}) bool {
	return internal.IsArray(a)

}

func IsSlice(a interface{}) bool {
	return internal.IsSlice(a)

}

func IsString(a interface{}) bool {
	return internal.IsString(a)

}

func IsBool(a interface{}) bool {
	return internal.IsBool(a)

}

func IsInteger(a interface{}) bool {
	return internal.IsInteger(a)

}

func IsFloat(a interface{}) bool {
	return internal.IsFloat(a)

}

func IsNumber(a interface{}) bool {
	return internal.IsNumber(a)

}

func IsFunc(a interface{}) bool {
	return internal.IsFunc(a)

}

func IsStruct(a interface{}) bool {
	return internal.IsStruct(a)

}

func IsPointer(a interface{}) bool {
	return internal.IsPointer(a)

}

func IsChan(a interface{}) bool {
	return internal.IsChan(a)

}

func IsMap(a interface{}) bool {
	return internal.IsMap(a)

}

func IsByteArray(a interface{}) bool {
	return internal.IsByteArray(a)

}

func IsEqual(a interface{}, b interface{}) bool {
	return internal.IsEqual(a, b)

}

func IsEmpty(a interface{}) bool {
	return IsEmpty(a)
}

func IsNil(a interface{}) bool {
	return IsNil(a)
}
