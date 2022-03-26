package main

import (
	"fmt"

	"github.com/gotility/gotil"
)

func main() {
	// a := gotil.IsAssigned("")

	var tt = [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}
	b := gotil.IsAssigned(tt)
	// fmt.Println(a)
	fmt.Println(b)
}
