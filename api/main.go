package main

import (
	"fmt"

	"github.com/gotilty/gotil/converter"
)

func main() {

	// var tt = [2]string{"ertugrul", "kutluer"}
	b, err := converter.ToInt64("ertugrul")
	// fmt.Println(a)
	fmt.Println(b, err)
}
