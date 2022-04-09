package gotil

import "github.com/gotilty/gotil/internal"

func Sort(a interface{}) (interface{}, error) {
	return SortBy(a, "")
}

func SortDesc(a interface{}) (interface{}, error) {
	return SortDescBy(a, "")
}

func SortBy(a interface{}, path string) (interface{}, error) {
	return internal.SortBy(a, path)
}

func SortDescBy(a interface{}, path string) (interface{}, error) {
	return internal.SortDescBy(a, path)
}
