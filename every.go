package gotil

// EveryBy checks if predicate returns truthy for each elements of given collection. Iteration is stopped once predicate returns false
func EveryBy[T any](s []T, f func(item T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// Contains checks if given item is exists in given collection. Iteration is stopped once predicate returns true
//	result := gotil.Contains([]int{5, 10, 15}, 10)
//	fmt.Println(result)
//	// Output: true
func Contains[T comparable](s []T, item T) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

// ContainsBy checks return true if given predicate return true.
func ContainsBy[T any](s []T, f func(item T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}
