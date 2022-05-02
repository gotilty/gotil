package gotil

// FilterBy iterates over elements of collection, returning an array of all elements predicate if returns true for.
// 	result := gotil.FilterBy([]int64{-100, -5, 30, 100}, func(v int64, i int) bool {
// 	return v > 0
// 	})
// 	fmt.Println(result)
// 	// Output: [30 100]
func FilterBy[T any](s []T, f func(v T, i int) bool) []T {
	result := []T{}
	for i, v := range s {
		if f(v, i) {
			result = append(result, v)
		}
	}
	return result
}
