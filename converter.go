package gotil

import (
	"github.com/gotilty/gotil/internal"
)

//ToFloat32 returns float32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToFloat32(a interface{}) (float32, error) {
	return internal.ToFloat32(a)
}

//ToFloat64 returns float64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToFloat64(a interface{}) (float64, error) {
	return internal.ToFloat64(a)

}

//ToInt returns int(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt(a interface{}) (int, error) {
	return internal.ToInt(a)

}

//ToInt8 returns int8(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt8(a interface{}) (int8, error) {
	return internal.ToInt8(a)

}

//ToInt16 returns int16(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt16(a interface{}) (int16, error) {
	return internal.ToInt16(a)

}

//ToInt32 returns int32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt32(a interface{}) (int32, error) {
	return internal.ToInt32(a)

}

//ToInt64 returns int64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToInt64(a interface{}) (int64, error) {
	return internal.ToInt64(a)

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
func ToString(a interface{}) (string, error) {
	return internal.ToString(a)

}

//Join just works with array and slice
//
//Example1:
//	Join([]int{10,20,30},"-") returns "10-20-30"
func Join(a interface{}, separator string) (string, error) {
	return internal.Join(a, separator)
}

//ToUint returns uint(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint(a interface{}) (uint, error) {
	return internal.ToUint(a)
}

//ToUint8 returns uint8(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint8(a interface{}) (uint8, error) {
	return internal.ToUint8(a)
}

//ToUint16 returns uint16(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint16(a interface{}) (uint16, error) {
	return internal.ToUint16(a)
}

//ToUint32 returns uint32(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint32(a interface{}) (uint32, error) {
	return internal.ToUint32(a)
}

//ToUint64 returns uint64(0) with an error if the parameter is unsupported type.
//Just works with all primitive types.
func ToUint64(a interface{}) (uint64, error) {
	return internal.ToUint64(a)
}
