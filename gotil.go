package gotil

import "fmt"

type Any interface {
	interface{}
}

// IsAssigned
func IsAssigned[K Any](a interface{}) bool {
	switch a.(type) {
	case int:
		// …
		fmt.Println("int")
	case string:
		fmt.Println("string")
		// …
	}
	return true
}
