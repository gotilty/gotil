package main

import (
	"fmt"

	"github.com/gotilty/gotil/collection"
)

func main() {
	m1 := []float64{5, 10.5, 10, 20, 10.75, 100, 4.23, 5.15, 5.99, 100.0001}

	result, _ := collection.Includes(m1, "gotil")
	fmt.Println(result)
}
