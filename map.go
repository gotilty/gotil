package gotil

// Creates an array of values by running each element in the given array thru iteratee.
// The value to be iterated should be given as the first parameter.
// The second parameter will be a function that will take the parameter and return value type interface{}.
// 	result := gotil.Map([]int{10, 20}, func(v, i int) int {
// 	return v * v
// 	})
// 	fmt.Println(result)
// 	// Output: 100, 400

func Map[T, D any](s []T, f func(val T, i int) D) []D {
	result := make([]D, len(s))
	for i, v := range s {
		result[i] = f(v, i)
	}
	return result
}
