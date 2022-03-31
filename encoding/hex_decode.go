package encoding

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"reflect"

	"github.com/gotilty/gotil/converter"
)

//HexDecode returns empty string with an error if the parameter is unsupported type.
//Just works with all primitive types, except arrays and slices.
//In decimal numbers, it should be used as a "string" in numbers where the total number of digits is too many.
func HexDecode(a interface{}) (string, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return hexToInt(val.Int())
	case reflect.Float32, reflect.Float64:
		return hexToFloat(val.Float())
	case reflect.String:
		return hexToString(val.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return hexToUint(val.Uint())
	default:
		return "", nil
	}
}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func hexToInt(s int64) (string, error) {
	return hexToString(converter.ToString(s))
}

func hexToFloat(s float64) (string, error) {
	return hexToString(converter.ToString(s))
}

func hexToUint(s uint64) (string, error) {
	return hexToString(converter.ToString(s))
}

// TODO: some errors need to be fixed
func hexToString(s string) (string, error) {
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x", dst[:n]), nil
}
