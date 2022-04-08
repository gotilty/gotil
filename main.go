package main

import (
	"fmt"
	"math"

	"github.com/gotilty/gotil/collection"
)

func main() {
	m1 := []float64{5, 10.5, 10, 10, 20, 20}
	if result, err := collection.GroupBy(m1, func(a interface{}, i int) interface{} {
		return math.Floor(a.(float64))
	}); err == nil {
		for _, v := range result.([]float64) {
			fmt.Println(v)
		}
		r := result.([]float64)
		fmt.Println(len(r))
	} else {

	}
}
