package encoding

import (
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/gotilty/gotil/converter"
)

//HexEncode returns empty string if the parameter is unsupported type.
//Just works with all primitive types, except arrays and slices.
//In decimal numbers, it should be used as a "string" in numbers where the total number of digits is too many.
func HexEncode(a interface{}) string {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return intToHex(val.Int())
	case reflect.Float32, reflect.Float64:
		return floatToHex(val.Float())
	case reflect.String:
		return stringToHex(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return uintToHex(val.Uint())
	default:
		return ""
	}
}

func intToHex(s int64) string {
	return stringToHex(converter.ToString(s))
}

func floatToHex(s float64) string {
	return stringToHex(converter.ToString(s))
}

func uintToHex(s uint64) string {
	return stringToHex(converter.ToString(s))
}

func stringToHex(s string) string {
	src := []byte(converter.ToString(s))
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return fmt.Sprintf("%s", dst)
}
