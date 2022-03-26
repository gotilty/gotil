package gotil

import "fmt"

func main() {
	var a int = 10
	var b string = ""
	fmt.Println("hello")
	t := IsAssigned[int](a)
	t2 := IsAssigned[string](b)

}
