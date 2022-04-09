package gotil

import "github.com/gotilty/gotil/internal"

func HexDecode(a interface{}) (string, error) {
	return internal.HexDecode(a)
}

func HexEncode(a interface{}) (string, error) {
	return internal.HexEncode(a)

}
