package main

import (
	"fmt"

	"github.com/gotilty/gotil/encoding"
)

func main() {

	// var tt = [2]string{"ertugrul", "kutluer"}
	//b, err := converter.ToInt64("ertugrul")
	c, _ := encoding.HexDecode("3543242.12314242")
	// fmt.Println(a)
	//fmt.Println(b, err)
	fmt.Println(c)
}
