package collection

import (
	"strings"

	"github.com/gotilty/gotil/converter"
)

func Includes(a interface{}, val interface{}) (bool, error) {
	a_str, _ := converter.ToString(a)
	val_str, _ := converter.ToString(val)
	return (strings.Contains(a_str, val_str)), nil
}
