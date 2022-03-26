package gotil

import (
	"fmt"
	"reflect"
)

type Any interface {
	interface{}
}

// IsAssigned
func IsAssigned[K Any](a K) bool {
	typeString := reflect.TypeOf(a).String()
	switch typeString {
	case "int":
		// …
		fmt.Println("int")
	case "string":
		fmt.Println("string")
		// …
	}
	return true
}
