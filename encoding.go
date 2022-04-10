package gotil

import "github.com/gotilty/gotil/internal"

//HexDecode returns empty string with an error if the parameter is unsupported type.
//Just works with all primitive types
//If given parameter is different type instead of string, firstly, it will convert the given parameter to string with gotil.ToString().
func HexDecode(a interface{}) (string, error) {
	return internal.HexDecode(a)
}

//HexEncode returns empty string if the parameter is unsupported type.
//Just works with all primitive types.
//In decimal numbers, it should be used as a "string" in numbers where the total number of digits is too many.
func HexEncode(a interface{}) (string, error) {
	return internal.HexEncode(a)

}
