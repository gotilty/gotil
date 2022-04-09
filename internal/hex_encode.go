package internal

import (
	"encoding/hex"
	"fmt"
)

//HexEncode returns empty string if the parameter is unsupported type.
//Just works with all primitive types, except arrays and slices.
//In decimal numbers, it should be used as a "string" in numbers where the total number of digits is too many.
func HexEncode(a interface{}) (string, error) {
	if s, err := ToString(a); err == nil {
		return stringToHex(s)
	} else {
		return "", err
	}
}
func stringToHex(s string) (string, error) {
	src := []byte(s)
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)
	return fmt.Sprintf("%s", dst), nil
}
