package array

import (
	"fmt"
	"reflect"
)

func Map(a interface{}, f func() interface{}) {
	fmt.Println(f())
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			fmt.Println(val.Index(i))
		}

	}
}
