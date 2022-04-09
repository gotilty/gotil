package internal

import (
	"reflect"

	"github.com/gotilty/gotil/internal/errs"
)

var map_type reflect.Type
var map_value reflect.Value
var value reflect.Value
var key reflect.Value
var result interface{}

func GroupBy(a interface{}, f func(val, i interface{}) interface{}) (interface{}, error) {
	val := reflect.ValueOf(a)
	kind := val.Kind()
	switch kind {
	case reflect.Slice, reflect.Array, reflect.Map:
		if kind == reflect.Slice || kind == reflect.Array {
			return groupBySlice(kind, val, f)
		}
		// return groupByMap(kind, val, f)
	}
	return nil, errs.NewUnsupportedTypeError(val.Kind().String())
}

func groupBySlice(kind reflect.Kind, val reflect.Value, f func(val, i interface{}) interface{}) (interface{}, error) {
	val_type := val.Type()
	assign := false
	var slice reflect.Value
	for k := 0; k < val.Len(); k++ {
		value = val.Index(k)
		result = f(value.Interface(), k)
		result_val := reflect.ValueOf(result)
		if assign == false {
			map_type = reflect.MapOf(
				reflect.TypeOf(result), val_type)
			map_value = reflect.MakeMap(map_type)
			slice = reflect.Zero(val_type)
			slice = reflect.Append(slice, value)
			map_value.SetMapIndex(result_val, slice)
			assign = true
		} else {
			keys := map_value.MapKeys()
			find_key, _ := FindBy(keys, func(val interface{}, i int) bool {
				key := val.(reflect.Value)
				if reflect.DeepEqual(key.Interface(), result_val.Interface()) {
					slice = map_value.MapIndex(key)
					return true
				}
				return false
			})
			if find_key != nil {
				slice = reflect.Append(slice, value)
				map_value.SetMapIndex(result_val, slice)
			} else {
				slice = reflect.Zero(val_type)
				slice = reflect.Append(slice, value)
				map_value.SetMapIndex(result_val, slice)
			}
		}
	}
	return map_value.Interface(), nil
}

// func groupByMap(kind reflect.Kind, val reflect.Value, f func(val, i interface{}) interface{}) (interface{}, error) {
// 	assign := false
// 	var slice reflect.Value
// 	keys := val.MapKeys()
// 	for k := 0; k < len(keys); k++ {
// 		key = keys[k]
// 		value = val.MapIndex(keys[k])
// 		result = f(value.Interface(), key.Interface())
// 		result_val := reflect.ValueOf(result)
// 		if assign == false {
// 			map_type = reflect.MapOf(reflect.TypeOf(result), reflect.SliceOf(reflect.TypeOf(value.Interface()).Elem()))
// 			map_value = reflect.MakeMap(map_type)
// 			slice = reflect.Zero(reflect.SliceOf(reflect.TypeOf(value.Interface()).Elem()))
// 			for i := 0; i < value.Len(); i++ {
// 				slice = reflect.Append(slice, value.Index(i))
// 			}
// 			map_value.SetMapIndex(result_val, slice)
// 			assign = true
// 		} else {
// 			keys := map_value.MapKeys()
// 			_find, _ := FindBy(keys, func(val interface{}, i int) bool {
// 				_val := val.(reflect.Value)
// 				if reflect.DeepEqual(_val.Interface(), result_val.Interface()) {
// 					slice = map_value.MapIndex(_val)
// 					return true
// 				}
// 				return false
// 			})
// 			if _find != nil {
// 				for i := 0; i < value.Len(); i++ {
// 					slice = reflect.Append(slice, value.Index(i))
// 				}

// 				map_value.SetMapIndex(result_val, slice)
// 			} else {
// 				slice = reflect.Zero(reflect.SliceOf(reflect.TypeOf(value.Interface()).Elem()))
// 				for i := 0; i < value.Len(); i++ {
// 					slice = reflect.Append(slice, value.Index(i))
// 				}
// 				map_value.SetMapIndex(result_val, slice)
// 			}
// 		}
// 	}
// 	return map_value.Interface(), nil
// }

// func addInMap(slice, result_val, map_value, value *reflect.Value, kind *reflect.Kind) reflect.Value {
// 	keys := map_value.MapKeys()
// 	_find, _ := FindBy(keys, func(val interface{}, i int) bool {
// 		_val := val.(reflect.Value)
// 		if reflect.DeepEqual(_val.Interface(), result_val.Interface()) {
// 			*slice = map_value.MapIndex(_val)
// 			return true
// 		}
// 		return false
// 	})

// 	if _find != nil {
// 		if *kind == reflect.Slice || *kind == reflect.Array {
// 			*slice = reflect.Append(*slice, *value)
// 			map_value.SetMapIndex(*result_val, *slice)
// 		} else {
// 			for i := 0; i < value.Len(); i++ {
// 				*slice = reflect.Append(*slice, value.Index(i))
// 			}
// 		}

// 		map_value.SetMapIndex(*result_val, *slice)
// 	} else {
// 		if *kind == reflect.Slice || *kind == reflect.Array {
// 			*slice = reflect.Zero((*value).Type())
// 			*slice = reflect.Append(*slice, *value)
// 			map_value.SetMapIndex(*result_val, *slice)
// 		} else {
// 			*slice = reflect.Zero(reflect.SliceOf(reflect.TypeOf((value).Interface()).Elem()))
// 			for i := 0; i < value.Len(); i++ {
// 				*slice = reflect.Append(*slice, value.Index(i))
// 			}
// 		}

// 		map_value.SetMapIndex(*result_val, *slice)
// 	}
// 	return *map_value
// }
