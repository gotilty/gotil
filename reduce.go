package gotil

// Reduce iterates given collection and returns the accumulated result of running each element
// the last param is initial value of accumulator.
//	result := Reduce([]int{5, 10}, func(accumulator, v, i int) int {
//	return accumulator + v
//	}, 0)
//	fmt.Println(result)
//	// Output: 15
func Reduce[T any, A any](s []T, f func(accumulator A, v T, i int) A, accumulator A) A {
	if len(s) < 1 {
		return accumulator
	}
	for i, v := range s {
		accumulator = f(accumulator, v, i)
	}
	return accumulator
}
