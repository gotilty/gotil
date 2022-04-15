package comp

import (
	"errors"
	"fmt"
	"strings"
)

func RunExample() {
	d1 := int64(10)
	d2 := int64(20)
	cp, err := getcomparer(d1)
	if err != nil {
		panic(err.Error())
	}
	result, err := cp.compare(d1, d2)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(result)
}

func getcomparer[T any](k T) (icomparable[T], error) {
	var comparer interface{}
	switch any(k).(type) {
	case int, int8, int16, int32, int64:
		comparer = integerComparer{}
		return comparer.(icomparable[T]), nil
	case string:
		comparer = stringComparer{}
		return (comparer).(icomparable[T]), nil
	default:
		return nil, errors.New("unsupported type")
	}
}

type icomparable[T any] interface {
	compare(a T, b T) (bool, error)
}

type stringComparer struct {
	icomparable[string]
}

type integerComparer struct {
	icomparable[int64]
}

func (s *stringComparer) compare(a *string, b *string) (bool, error) {
	return strings.Compare((*a), (*b)) == -1, nil
}

func (s *integerComparer) compare(a *int64, b *int64) (bool, error) {
	return (*a) < (*b), nil
}
