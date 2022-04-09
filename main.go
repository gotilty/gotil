package main

import (
	"fmt"

	"github.com/gotilty/gotil/converter"
)

func main() {
	b, _ := converter.ToUint16(65537)
	fmt.Println(b)
}
