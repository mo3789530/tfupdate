package utils

import "reflect"

func Type(t interface{}) reflect.Type {
	return reflect.TypeOf(t)
}
