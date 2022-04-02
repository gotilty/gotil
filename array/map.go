package array

import (
	"fmt"
	"reflect"

	"github.com/gotilty/gotil/converter"
)

func Map(a interface{}, f func(i interface{}) interface{}) interface{} {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		typ := reflect.TypeOf(a).Elem().Kind()
		if typ == reflect.Int || typ == reflect.Int8 || typ == reflect.Int16 || typ == reflect.Int32 || typ == reflect.Int64 {
			for i := 0; i < val.Len(); i++ {
				var res = []int64{}
				vall := val.Index(i).Interface()
				int_value, _ := converter.ToInt64(vall.(int))
				fmt.Println(f(int_value))
				// ress, _ := converter.ToInt64(f(int_value).(int))
				a = append(res, int_value)
			}
		}
	}
	return a
}
