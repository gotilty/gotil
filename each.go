package gotil

// Creates an array of values by running each element in the given array thru iteratee.
// The value to be iterated should be given as the first parameter.
// The second parameter will be a function that will take the parameter and return value type interface{}.
// 	gotil.Each([]string{"gotilty", "gotil"}, func(v string) {
// 	fmt.Fprint(os.Stdout, v)
// 	})
// 	// Output: gotiltygotil

func Each[T any](s []T, f func(v T)) {
	for _, v := range s {
		f(v)
	}
}
