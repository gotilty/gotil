package internal

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gotilty/gotil/internal/errs"
)

type sortResult struct {
	r   interface{}
	p   string
	e   error
	asc bool
}

func (s sortResult) getResult() (interface{}, error) {
	return s.r, s.e
}

// Sort returns a new array of ordered values as ascending, using a version of Merge Algorithm.
//
// data := []int64{100, 30, -100, -5}
// newData, _ := Sort(data)
//
// Output: [-100 -5 30 100]
func Sort(a interface{}) (interface{}, error) {
	return SortBy(a, "")
}

// SortDesc returns a new array of ordered values as descending, using a version of Merge Algorithm.
//
// data := []int64{100, 30, -100, -5}
// newData, _ := SortDesc(data)
//
// Output: [100 30 -5 -100]
func SortDesc(a interface{}) (interface{}, error) {
	return SortDescBy(a, "")
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
	sr := sortResult{
		p:   path,
		asc: true,
	}
	sr = sr.sortByThen(a, path)
	return sr.getResult()
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
	sr := sortResult{
		p:   path,
		asc: false,
	}
	sr = sr.sortByThen(a, path)
	return sr.getResult()
}

func (s sortResult) sortByThen(a interface{}, path string) sortResult {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		newSlice := copySlice(a)
		if newSlice.Len() < 2 {
			r := newSlice.Interface()
			s.r = r
			s.e = nil
			return s
		}
		k := newSlice.Index(0).Type().Kind()
		c, errc := getcomparer(k, path)
		if errc != nil {
			s.r = nil
			s.e = errc
			return s
		}
		if mr, errm := mergeSort(&newSlice, c, s.asc); errm == nil {
			r := mr.Interface()
			s.r = r
			s.e = nil
			return s
		} else {
			s.r = nil
			s.e = errm
			return s
		}
	}
	s.r = nil
	s.e = errs.NewUnsupportedTypeError(val.Kind().String())
	return s
}

func mergeSort(val *reflect.Value, c icomparable, asc bool) (*reflect.Value, error) {
	if (*val).Len() < 2 {
		return (val), nil
	}
	li := (*val).Len() / 2
	left := (*val).Slice(0, li)
	first, errf := mergeSort(&left, c, asc)
	if errf != nil {
		return nil, errf
	}
	right := (*val).Slice(li, (*val).Len())
	second, errs := mergeSort(&right, c, asc)
	if errs != nil {
		return nil, errs
	}
	return merge(first, second, c, asc)
}

func merge(a *reflect.Value, b *reflect.Value, c icomparable, asc bool) (*reflect.Value, error) {
	t := reflect.TypeOf((*a).Interface())
	final := reflect.MakeSlice(t, 0, (*a).Len()+(*b).Len())

	i := 0
	j := 0

	for i < (*a).Len() && j < (*b).Len() {
		ii := (*a).Index(i)
		ij := (*b).Index(j)
		if ok, err := c.compare(&ii, &ij); err == nil {
			if asc == false {
				ok = !ok
			}
			if ok {
				final = reflect.Append(final, reflect.ValueOf(ii.Interface()))
				i++
			} else {
				final = reflect.Append(final, reflect.ValueOf(ij.Interface()))
				j++
			}
		} else {
			return nil, err
		}
	}
	for ; i < (*a).Len(); i++ {
		ii := (*a).Index(i)
		final = reflect.Append(final, ii)
	}

	for ; j < (*b).Len(); j++ {
		ij := (*b).Index(j)
		final = reflect.Append(final, ij)
	}
	return &final, nil
}

func getcomparer(k reflect.Kind, path string) (icomparable, error) {
	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &integerComparer{}, nil
	case reflect.Float32, reflect.Float64:
		return &floatComparer{}, nil
	case reflect.String:
		return &stringComparer{}, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return &integerComparer{}, nil
	case reflect.Bool:
		return &boolComparer{}, nil
	case reflect.Struct:
		if path == "" {
			return nil, errs.NewMissingParameterError("path", "")
		}
		properties := strings.SplitN(path, ".", 2)
		property := ""
		if len(properties) >= 2 {
			property = properties[0]
			path = properties[1]
		} else {
			property = properties[0]
			path = ""
		}
		return &structComparer{
			property: property,
			path:     path,
		}, nil
	case reflect.Array, reflect.Slice:
		return &sliceComparer{}, nil
	default:
		return nil, errs.NewUnsupportedTypeError(k.String())
	}
}

// for i := newSlice.Len() - 1; i > 0; i-- { // Fisher–Yates shuffle
// 	j := rand.Intn(i + 1)
// 	tp := newSlice.Index(i).Interface()
// 	newSlice.Index(i).Set(newSlice.Index(j))
// 	newSlice.Index(j).Set(reflect.ValueOf(tp))
// }

type icomparable interface {
	compare(a *reflect.Value, b *reflect.Value) (bool, error)
}

type stringComparer struct {
	icomparable
}

type integerComparer struct {
	icomparable
}

type floatComparer struct {
	icomparable
}

type boolComparer struct {
	icomparable
}

type structComparer struct {
	icomparable,
	property string
	path  string
	child icomparable
}

type sliceComparer struct {
	icomparable
}

func (s *stringComparer) compare(a *reflect.Value, b *reflect.Value) (bool, error) {
	return strings.Compare((*a).String(), (*b).String()) == -1, nil
}

func (s *integerComparer) compare(a *reflect.Value, b *reflect.Value) (bool, error) {
	return (*a).Int() < (*b).Int(), nil
}

func (s *floatComparer) compare(a *reflect.Value, b *reflect.Value) (bool, error) {
	return (*a).Float() < (*b).Float(), nil
}

func (s *boolComparer) compare(a *reflect.Value, b *reflect.Value) (bool, error) {
	return (*a).Bool() != true, nil
}

func (s *structComparer) compare(a *reflect.Value, b *reflect.Value) (bool, error) {
	ap := (*a).FieldByName(s.property)
	bp := (*b).FieldByName(s.property)
	apk := ap.Kind()
	bpk := bp.Kind()
	if apk.String() != bpk.String() {
		return false, errs.NewUnsupportedParameterTypeError(
			fmt.Sprintf("%s [%v]: %s [%v]",
				apk.String(),
				ap.Interface(),
				bpk.String(),
				bp.Interface()),
			"they must be the same type",
		)
	}
	if checkKind(apk, reflect.Struct) {
		if s.path == "" {
			return false, errs.NewMissingParameterError(apk.String(), fmt.Sprintf("set a value in this property %s", s.property))
		} else {
			if cn, err := getcomparer(apk, s.path); err != nil {
				return false, err
			} else {
				return cn.compare(&ap, &bp)
			}
		}
	} else {
		if cn, err := getcomparer(apk, ""); err != nil {
			return false, err
		} else {
			return cn.compare(&ap, &bp)
		}
	}
}

func (s *sliceComparer) compare(a *reflect.Value, b *reflect.Value) (bool, error) {
	//TODO: slice compare
	return (*a).Bool() != true, nil
}
