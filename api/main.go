package main

import (
	"fmt"

	"github.com/gotility/gotil"
)

func main() {
	// a := gotil.IsAssigned("")

	var tt = [0]string{}
	b := gotil.IsAssigned(tt)
	// fmt.Println(a)
	fmt.Println(b)
}
