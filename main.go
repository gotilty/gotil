package main

import (
	"github.com/gotilty/gotil/array"
)

func main() {
	g := greetings
	array.Map("ertu", g)
}

func greetings() interface{} {
	return ("ertu")
}
