package internal

import (
	"strings"
)

func Includes(a interface{}, val interface{}) (bool, error) {
	a_str, _ := ToString(a)
	val_str, _ := ToString(val)
	return (strings.Contains(a_str, val_str)), nil
}
