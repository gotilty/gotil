package main

import (
	"fmt"

	"github.com/gotilty/gotil/array"
	"github.com/gotilty/gotil/converter"
)

func main() {
	m1 := []int64{5, 10}
	if result, err := array.Map(m1, func(a interface{}, i int) interface{} {
		a1, _ := converter.ToString(a)
		_, _ = converter.ToInt64(i)
		return a1 + "ertu"
	}); err == nil {
		for _, v := range result.([]string) {
			fmt.Println(v)
		}
	} else {

	}
}
