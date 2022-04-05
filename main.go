package main

import (
	"fmt"

	"github.com/gotilty/gotil/array"
	"github.com/gotilty/gotil/converter"
)

func main() {
	// sample_array1 := []uint64{9, 12, 15, 34525423452435234, 4, 4532584375937459345, 3, 1, 345345342342423523, 21, 12}
	sample_array2 := []string{"Gotilty", "Gotil"}
	// a, _ := array.Map(sample_array1, square)
	// b := a.([]uint64)
	sample_res2, _ := array.Map(sample_array2, addChar)
	fmt.Println(sample_res2.([]string))
	// for _, a := range sample_res2 {
	// 	fmt.Println(a)
	// }
}

func square(a interface{}, i int) interface{} {
	b, _ := converter.ToUint64(a)
	return b * 2
}

func addChar(a interface{}, i int) interface{} {
	b, _ := converter.ToString(a)
	if i == 0 {
		return b + "->"
	}
	return b + " is the best utility tool for Golang!"
}
