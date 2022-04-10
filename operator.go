package gotil

import "github.com/gotilty/gotil/internal"

// IsAssigned checks the given parameter has value also if it's slice or array that checks has element.
//
// Example:
//
// result := IsAssigned("test data")
//
// Output: true
func IsAssigned(a interface{}) bool {
	return internal.IsAssigned(a)
}

// IsNotAssigned checks the given parameter has not value also if it's slice or array that checks has not any element.
// It works same IsEmpty()
//
// Example:
//
// result := IsNotAssigned("test data")
//
// Output: false
func IsNotAssigned(a interface{}) bool {
	return internal.IsNotAssigned(a)
}

// IsArray checks the given parameter is an array
func IsArray(a interface{}) bool {
	return internal.IsArray(a)

}

// IsSlice checks the given parameter is a slice
func IsSlice(a interface{}) bool {
	return internal.IsSlice(a)

}

// IsString checks the given parameter is a string
func IsString(a interface{}) bool {
	return internal.IsString(a)

}

// IsBool checks the given parameter is a bool
func IsBool(a interface{}) bool {
	return internal.IsBool(a)

}

// IsInteger checks the given parameter is an integer
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, unit64
func IsInteger(a interface{}) bool {
	return internal.IsInteger(a)

}

// IsFloat checks the given parameter is an float
// float32, float64
func IsFloat(a interface{}) bool {
	return internal.IsFloat(a)

}

// IsNumber checks the given parameter is type of any golang number
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, unit64, float32, float64
func IsNumber(a interface{}) bool {
	return internal.IsNumber(a)

}

// IsFunc checks the given parameter is a func
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, unit64, float32, float64
func IsFunc(a interface{}) bool {
	return internal.IsFunc(a)

}

// IsStruct checks the given parameter is a struct
func IsStruct(a interface{}) bool {
	return internal.IsStruct(a)

}

// IsPointer checks the given parameter is a pointer
func IsPointer(a interface{}) bool {
	return internal.IsPointer(a)

}

// IsChan checks the given parameter is a channel
func IsChan(a interface{}) bool {
	return internal.IsChan(a)

}

// IsMap checks the given parameter is a map
func IsMap(a interface{}) bool {
	return internal.IsMap(a)

}

// IsByteArray checks the given parameter is a byte array
func IsByteArray(a interface{}) bool {
	return internal.IsByteArray(a)

}

// IsEqual checks the given any 2 parameters are equal
func IsEqual(a interface{}, b interface{}) bool {
	return internal.IsEqual(a, b)

}

// IsEmpty checks the given parameter has not value also if it's slice or array that checks has not any element.
// It works same IsNotAssigned()
func IsEmpty(a interface{}, b interface{}) bool {
	return internal.IsEmpty(a, b)
}

// IsNil checks the given parameter has not value also if it's slice or array that checks has not any element.
// It works same IsNotAssigned()
func IsNil(a interface{}, b interface{}) bool {
	return internal.IsNil(a, b)
}

// Size returns a length of given parameter
//
// data := []int64{100, 30, -100, -5}
// result, _ := Size(data)
// result2, _ := Size("gotil")

// Output: 4
// Output2: 5
func Size(val interface{}) (int, error) {
	return internal.Size(val)
}
