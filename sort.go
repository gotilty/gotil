package gotil

import "github.com/gotilty/gotil/internal"

// Sort returns a new array of ordered values as ascending, using a version of Merge Algorithm.
//
// data := []int64{100, 30, -100, -5}
// newData, _ := Sort(data)
//
// Output: [-100 -5 30 100]
func Sort(a interface{}) (interface{}, error) {
	return internal.SortBy(a, "")
}

// SortDesc returns a new array of ordered values as descending, using a version of Merge Algorithm.
//
// data := []int64{100, 30, -100, -5}
// newData, _ := SortDesc(data)
//
// Output: [100 30 -5 -100]
func SortDesc(a interface{}) (interface{}, error) {
	return internal.SortDescBy(a, "")
}

// SortBy returns a new array of ordered values as ascending, using a version of Merge Algorithm with a given property path list of struct
//
// Example: SortBy(data, "location.city")
//
// data := []user{
// 	{
// 		name: "Micheal",
// 		age:  27,
// 		location: location{
// 			city: "New York",
// 		},
// 	},
// 	{
// 		name: "Joe",
// 		age:  30,
// 		location: location{
// 			city: "Detroit",
// 		},
// 	},
// }
//
// newData, _ := SortBy(data, "location.city")
//
// Output: [{Joe 30 {Detroit}} {Micheal 27 {New York}}]
func SortBy(a interface{}, path string) (interface{}, error) {
	return internal.SortBy(a, path)
}

// SortDescBy returns a new array of ordered values as descending, using a version of Merge Algorithm with a given property path list of struct
//
// Example: SortDescBy(data, "location.city")
//
// data := []user{
// 	{
// 		name: "Micheal",
// 		age:  27,
// 		location: location{
// 			city: "New York",
// 		},
// 	},
// 	{
// 		name: "Joe",
// 		age:  30,
// 		location: location{
// 			city: "Detroit",
// 		},
// 	},
// }
//
// newData, _ := SortDescBy(data, "location.city")
//
// Output: [{Micheal 27 {New York}} {Joe 30 {Detroit}}]
func SortDescBy(a interface{}, path string) (interface{}, error) {
	return internal.SortDescBy(a, path)
}
