package main

import (
	"fmt"
	"strings"

	"github.com/gotilty/gotil/collection"
	"github.com/gotilty/gotil/converter"
)

func main() {
	m1 := map[int][]string{1: {"gotilty"}, 2: {"gotil"}, 3: {"is"}, 4: {"perfect"}, 5: {"and"}, 6: {"easy"}, 7: {"to"}, 8: {"use"}}

	result, _ := collection.GroupBy(m1, func(val interface{}, i interface{}) interface{} {
		for j := 0; j < len(val.([]string)); j++ {
			value, _ := converter.ToString(val)
			return len(strings.Split(value, ""))
		}
		return len(strings.Split(val.(string), ""))
	})
	fmt.Println(result)
}
