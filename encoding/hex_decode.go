package encoding

import (
	"encoding/hex"
	"fmt"

	"github.com/gotilty/gotil/converter"
)

//HexDecode returns empty string with an error if the parameter is unsupported type.
//Just works with all primitive types, except arrays and slices.
//If given parameter is different type instead of string, firstly, it will convert the given parameter to string.
func HexDecode(a interface{}) (string, error) {
	if s, err := converter.ToString(a); err == nil {
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
