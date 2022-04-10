package internal

import (
	"encoding/hex"
	"fmt"
)

//HexDecode returns empty string with an error if the parameter is unsupported type.
//Just works with all primitive types
//If given parameter is different type instead of string, firstly, it will convert the given parameter to string with gotil.ToString().
func HexDecode(a interface{}) (string, error) {
	if s, err := ToString(a); err == nil {
		return hexToString(s)
	} else {
		return "", err
	}
}

func hexToString(s string) (string, error) {
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	if n, err := hex.Decode(dst, src); err == nil {
		return fmt.Sprintf("%s", dst[:n]), nil
	} else {
		return "", err
	}
}
