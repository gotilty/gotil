package gotil

import (
	"fmt"
	"reflect"
)

type Base struct {
	int64
	int32
}

// If source is string. Fastest and most reliable way to check if string is empty is s != ""
func isAssignedStr(s string) bool {
	return s != ""
}
func isAssignedInt(s Base) bool {
	return s.String() != 0
}

// IsAssigned
func IsAssigned[K interface{}](a K) bool {
	typeString := reflect.TypeOf(a).String()
	switch typeString {
	case "int":
		fmt.Println("int")
	case "string":
		return isAssignedStr(fmt.Sprintf("%v", a))
	}
	return true
}
