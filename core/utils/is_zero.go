package utils

import "reflect"

func IsZero(val interface{}) bool {
	v := reflect.ValueOf(val)
	return !v.IsValid() || reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
