package gotil

// GroupBy iterates over elements of collection, returns an map object.
// key is result of iterator method
// value is list
// 	result := gotil.GroupBy([]int{10, 20, 30}, func(v int, i int) int {
// 	return v / 10
// 	})
// 	fmt.Println(result)
// 	// Output: map[1:[10] 2:[20] 3:[30]]
func GroupBy[T, K comparable](s []T, f func(v T, i int) K) map[K][]T {
	result := map[K][]T{}
	for i, v := range s {
		key := f(v, i)
		result[key] = append(result[key], v)
	}
	return result
}
