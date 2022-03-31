package main

import (
	"fmt"

	"github.com/gotilty/gotil/converter"
)

func main() {

	// var tt = [2]string{"ertugrul", "kutluer"}
	b, _ := converter.ToString(123123123.1231233)
	// fmt.Println(a)
	fmt.Println(b)
}
