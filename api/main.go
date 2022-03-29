package main

import (
	"fmt"

	"github.com/gotility/gotil"
)

func main() {

	var tt = [2]string{"ertugrul", "kutluer"}
	b := gotil.ToString(tt, ",")
	// fmt.Println(a)
	fmt.Println(b)
}
