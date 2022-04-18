package gotil

// FindBy iterates over elements of collection, returns an object of fist element predicate if returns true for.
//	data := []int64{-100, -5, 30, 100}
//	newData, _ := FindBy(data, func(val interface{}, i int) bool {
//	if val.(int64) == 30 {
// 			return true
// 		} else {
// 			return false
// 		}
// 	})
// 	// Output: 30
func FindBy[T any](s []T, f func(val T) bool) T {
	return FindByFromIndex(s, f, 0)
}

// FindByFromIndex iterates over elements of collection, returns an object of fist element predicate if returns true for.
// It works same as FindBy just from parameter provides to set start index of given slice or array.
func FindByFromIndex[T any](s []T, f func(val T) bool, from int) T {
	var result T
	index := findIndex(s, f, 0)
	if index > -1 {
		return s[index]
	}
	return result
}

// FindLastBy iterates over elements of collection, returns an object of lastest element predicate if returns true for.
// 	data := []int64{-100, -5, 30, 100}
// 	newData, _ := FindLastBy(data, func(val interface{}, i int) bool {
// 		if val.(int64) == 30 {
// 			return true
// 		} else {
// 			return false
// 		}
// 	})
// 	// Output: 30
func FindLastBy[T any](s []T, f func(val T) bool) T {
	return FindLastByFromIndex(s, f, 0)
}

// FindLastBy iterates over elements of collection, returns an object of lastest element predicate if returns true for.
// It works same as FindLastBy just from parameter provides to set start index of given slice or array.
func FindLastByFromIndex[T any](s []T, f func(val T) bool, from int) T {
	var result T
	index := findLastIndex(s, f, 0)
	if index > -1 {
		return s[index]
	}
	return result
}

func findIndex[T any](s []T, f func(val T) bool, from int) int {
	length := len(s)
	if length == 0 {
		return -1
	}
	index := 0
	if from < 0 {
		index = max(length+index, 0)
	} else {
		index = min(index, length-1)
	}
	return baseFindIndex(s, f, index, false)
}

func findLastIndex[T any](s []T, f func(val T) bool, from int) int {
	length := len(s)
	if length == 0 {
		return -1
	}
	index := length - 1
	if from < 0 {
		index = max(length+index, 0)
	} else {
		index = min(index, length-1)
	}
	return baseFindIndex(s, f, index, true)
}

func baseFindIndex[T any](s []T, f func(val T) bool, from int, right bool) int {
	length := len(s)
	index := from
	if right {
		for index < length {
			value := s[index]
			if f(value) {
				return index
			}
			index--
		}
	} else {
		for index < length {
			value := s[index]
			if f(value) {
				return index
			}
			index++
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
