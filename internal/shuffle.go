package internal

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/gotilty/gotil/internal/errs"
)

func Shuffle(a interface{}) (interface{}, error) {
	return ShuffleSeed(a, time.Now().UnixNano())
}

func ShuffleSeed(a interface{}, seed int64) (interface{}, error) {
	val := reflect.ValueOf(a)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		rand.Seed(seed)
		newSlice := copySlice(a)
		for i := newSlice.Len() - 1; i > 0; i-- { // Fisher–Yates shuffle
			j := rand.Intn(i + 1)
			tp := newSlice.Index(i).Interface()
			newSlice.Index(i).Set(newSlice.Index(j))
			newSlice.Index(j).Set(reflect.ValueOf(tp))
		}
		return newSlice.Interface(), nil
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}

func copySlice(src interface{}) reflect.Value {
	t := reflect.TypeOf(src)
	if t.Kind() != reflect.Slice {
		panic("not a slice!")
	}
	v := reflect.ValueOf(src)

	target := reflect.MakeSlice(t, v.Cap(), v.Len())
	reflect.Copy(target, v)
	return target
}
