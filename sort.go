package gotil

import (
	"errors"
	"strings"
)

// Sort returns a new array of ordered values as ascending, using a version of Merge Algorithm.
//	data := []int64{100, 30, -100, -5}
//	newData, _ := gotil.Sort(data)
//	// Output: [-100 -5 30 100]
func Sort[T any](s []T) []T {
	return SortBy(s, "")
}

// SortDesc returns a new array of ordered values as descending, using a version of Merge Algorithm.
//	data := []int64{100, 30, -100, -5}
//	newData, _ := SortDesc(data)
//	// Output: [100 30 -5 -100]
func SortDesc[T any](s []T) []T {
	return SortDescBy(s, "")
}

// SortBy returns a new array of ordered values as ascending, using a version of Merge Algorithm with a given property path list of struct
//
// Example:
//	SortBy(data, "location.city")
//	data := []user{
//		{
//			name: "Micheal",
// 			age:  27,
// 			location: location{
// 				city: "New York",
// 			},
//		},
//		{
// 			name: "Joe",
// 			age:  30,
// 			location: location{
// 				city: "Detroit",
// 			},
//  	},
//	}
//	newData, _ := SortBy(data, "location.city")
//	// Output: [{Joe 30 {Detroit}} {Micheal 27 {New York}}]
func SortBy[T any](s []T, path string) []T {
	sr := sortResult[T]{
		p:   path,
		asc: true,
	}
	sr = sr.sortByThen(s, path)
	r, _ := sr.getResult()
	return r
}

// SortDescBy returns a new array of ordered values as descending, using a version of Merge Algorithm with a given property path list of struct
//
// Example: SortDescBy(data, "location.city")
//	data := []user{
//  	{
// 			name: "Micheal",
// 			age:  27,
// 			location: location{
// 				city: "New York",
// 			},
// 		},
//		{
// 			name: "Joe",
// 			age:  30,
// 			location: location{
// 				city: "Detroit",
// 			},
// 		 },
//	}
//	newData, _ := SortDescBy(data, "location.city")
//	// Output: [{Micheal 27 {New York}} {Joe 30 {Detroit}}]
func SortDescBy[T any](s []T, path string) []T {
	sr := sortResult[T]{
		p:   path,
		asc: false,
	}
	sr = sr.sortByThen(s, path)
	r, _ := sr.getResult()
	return r
}

type sortResult[T any] struct {
	r   []T
	p   string
	e   error
	asc bool
}

func (s sortResult[T]) getResult() ([]T, error) {
	return s.r, s.e
}

func (sr sortResult[T]) sortByThen(s []T, path string) sortResult[T] {
	newSlice := copySlice(s)
	if len(newSlice) < 2 {
		sr.r = newSlice
		sr.e = nil
		return sr
	}
	k := newSlice[0]
	c, errc := getcomparer(&k, "")
	if errc != nil {
		sr.r = nil
		sr.e = errc
		return sr
	}
	if mr, errm := mergeSort(&newSlice, c, sr.asc); errm == nil {
		sr.r = (*mr)
		sr.e = nil
		return sr
	} else {
		sr.r = nil
		sr.e = errm
		return sr
	}
}

func mergeSort[T any](val *[]T, c icomparable[*T], asc bool) (*[]T, error) {
	if len(*val) < 2 {
		return (val), nil
	}
	li := len(*val) / 2
	left := (*val)[0:li]
	first, errf := mergeSort(&left, c, asc)
	if errf != nil {
		return nil, errf
	}
	right := (*val)[li:len(*val)]
	second, errs := mergeSort(&right, c, asc)
	if errs != nil {
		return nil, errs
	}
	return merge(first, second, c, asc)
}

func merge[T any](a *[]T, b *[]T, c icomparable[*T], asc bool) (*[]T, error) {
	final := make([]T, 0, len(*a)+len(*b))
	i := 0
	j := 0

	for i < len(*a) && j < len(*b) {
		ii := (*a)[i]
		ij := (*b)[j]
		if ok, err := c.compare(&ii, &ij); err == nil {
			if asc == false {
				ok = !ok
			}
			if ok {
				final = append(final, ii)
				i++
			} else {
				final = append(final, ij)
				j++
			}
		} else {
			return nil, err
		}
	}
	for ; i < len(*a); i++ {
		ii := (*a)[i]
		final = append(final, ii)
	}

	for ; j < len(*b); j++ {
		ij := (*b)[j]
		final = append(final, ij)
	}
	return &final, nil
}
func getcomparer[T any](k T, path string) (icomparable[T], error) {
	var comparer interface{}
	switch any(k).(type) {
	case *int:
		comparer = intComparer{}
		return comparer.(icomparable[T]), nil
	case *int8:
		comparer = int8Comparer{}
		return comparer.(icomparable[T]), nil
	case *int16:
		comparer = int16Comparer{}
		return comparer.(icomparable[T]), nil
	case *int32:
		comparer = int32Comparer{}
		return comparer.(icomparable[T]), nil
	case *int64:
		comparer = int64Comparer{}
		return comparer.(icomparable[T]), nil
	case *uint:
		comparer = uIntComparer{}
		return comparer.(icomparable[T]), nil
	case *uint8:
		comparer = uInt8Comparer{}
		return comparer.(icomparable[T]), nil
	case *uint16:
		comparer = uInt16Comparer{}
		return comparer.(icomparable[T]), nil
	case *uint32:
		comparer = uInt32Comparer{}
		return comparer.(icomparable[T]), nil
	case *uint64:
		comparer = uInt64Comparer{}
		return comparer.(icomparable[T]), nil
	case *string:
		comparer = stringComparer{}
		return (comparer).(icomparable[T]), nil
	default:

		//TODO: impl struct
		// if path == "" {
		// 	return nil, errs.NewMissingParameterError("path", "")
		// }
		// properties := strings.SplitN(path, ".", 2)
		// property := ""
		// if len(properties) >= 2 {
		// 	property = properties[0]
		// 	path = properties[1]
		// } else {
		// 	property = properties[0]
		// 	path = ""
		// }
		// comparer = structComparer[T]{
		// 	property: property,
		// 	path:     path,
		// }
		// return (comparer).(icomparable[T]), nil
		return nil, errors.New("unsupported type")
	}
}

func copySlice[T any](src []T) []T {
	new := make([]T, len(src))
	copy(new, src)
	return new
}

type icomparable[T any] interface {
	compare(a T, b T) (bool, error)
}

type stringComparer struct {
	icomparable[*string]
}

type intComparer struct {
	icomparable[*int]
}

type int8Comparer struct {
	icomparable[*int8]
}
type int16Comparer struct {
	icomparable[*int16]
}

type int32Comparer struct {
	icomparable[*int32]
}

type int64Comparer struct {
	icomparable[*int64]
}

type uIntComparer struct {
	icomparable[*uint]
}

type uInt8Comparer struct {
	icomparable[*uint8]
}
type uInt16Comparer struct {
	icomparable[*uint16]
}

type uInt32Comparer struct {
	icomparable[*uint32]
}

type uInt64Comparer struct {
	icomparable[*uint64]
}

type float32Comparer struct {
	icomparable[*float32]
}
type float64Comparer struct {
	icomparable[*float64]
}

type boolComparer struct {
	icomparable[*bool]
}

type Integer interface {
	int | int8 | int16 | int32 | int64
}

type structComparer[T any] struct {
	icomparable,
	property string
	path  string
	child icomparable[T]
}

func (s stringComparer) compare(a *string, b *string) (bool, error) {
	return strings.Compare((*a), (*b)) == -1, nil
}

func (s intComparer) compare(a *int, b *int) (bool, error) {
	return (*a) < (*b), nil
}
func (s int8Comparer) compare(a *int8, b *int8) (bool, error) {
	return (*a) < (*b), nil
}
func (s int16Comparer) compare(a *int16, b *int16) (bool, error) {
	return (*a) < (*b), nil
}

func (s int32Comparer) compare(a *int32, b *int32) (bool, error) {
	return (*a) < (*b), nil
}
func (s int64Comparer) compare(a *int64, b *int64) (bool, error) {
	return (*a) < (*b), nil
}

func (s uIntComparer) compare(a *uint, b *uint) (bool, error) {
	return (*a) < (*b), nil
}
func (s uInt8Comparer) compare(a *uint8, b *uint8) (bool, error) {
	return (*a) < (*b), nil
}
func (s uInt16Comparer) compare(a *uint16, b *uint16) (bool, error) {
	return (*a) < (*b), nil
}

func (s uInt32Comparer) compare(a *uint32, b *uint32) (bool, error) {
	return (*a) < (*b), nil
}
func (s uInt64Comparer) compare(a *uint64, b *uint64) (bool, error) {
	return (*a) < (*b), nil
}

func (s float32Comparer) compare(a *float32, b *float32) (bool, error) {
	return (*a) < (*b), nil
}

func (s float64Comparer) compare(a *float64, b *float64) (bool, error) {
	return (*a) < (*b), nil
}

func (s boolComparer) compare(a *bool, b *bool) (bool, error) {
	return (*a) != true, nil
}

// func (s *structComparer[T]) compare(a *T, b *T) (bool, error) {
// 	ap := (*a).FieldByName(s.property)
// 	bp := (*b).FieldByName(s.property)
// 	apk := ap.Kind()
// 	bpk := bp.Kind()
// 	if apk.String() != bpk.String() {
// 		return false, errs.NewUnsupportedParameterTypeError(
// 			fmt.Sprintf("%s [%v]: %s [%v]",
// 				apk.String(),
// 				ap.Interface(),
// 				bpk.String(),
// 				bp.Interface()),
// 			"they must be the same type",
// 		)
// 	}
// 	if checkKind(apk, reflect.Struct) {
// 		if s.path == "" {
// 			return false, errs.NewMissingParameterError(apk.String(), fmt.Sprintf("set a value in this property %s", s.property))
// 		} else {
// 			if cn, err := getcomparer(apk, s.path); err != nil {
// 				return false, err
// 			} else {
// 				return cn.compare(&ap, &bp)
// 			}
// 		}
// 	} else {
// 		if cn, err := getcomparer(apk, ""); err != nil {
// 			return false, err
// 		} else {
// 			return cn.compare(&ap, &bp)
// 		}
// 	}
// }
