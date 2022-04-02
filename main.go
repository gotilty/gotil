package main

import (
	"fmt"

	"github.com/gotilty/gotil/array"
	"github.com/gotilty/gotil/converter"
)

func main() {
	s := square
	sample_array1 := []int{9, -13, 4, -2, 3, 1, -10, 21, 12}
	fmt.Println(array.Map(sample_array1, s))
}

func square(a interface{}) interface{} {
	b, _ := converter.ToInt64(a)
	return (b * 3)
}
