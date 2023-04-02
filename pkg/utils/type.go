package utils

import (
	"fmt"
	"reflect"
)

func TypeString(t interface{}) string {
	return fmt.Sprintf("%v", reflect.TypeOf(t))
}
