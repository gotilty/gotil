package main

import (
	"fmt"
	"strconv"

	"github.com/gotility/gotil"
)

func main() {

	str1 := "123"

	// using ParseInt method
	int1, err := strconv.ParseInt(str1, 6, 12)

	fmt.Println(int1)

	// If the input string contains the integer
	// greater than base, get the zero value in return
	str2 := "123"

	int2, err := strconv.ParseInt(str2, 2, 32)
	fmt.Println(int2, err)

	// a := gotil.IsAssigned("")

	var tt = [2]string{"ertugrul", "kutluer"}
	b := gotil.ToString(tt)
	// fmt.Println(a)
	fmt.Println(b)
}
