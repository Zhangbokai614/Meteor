package utils

import (
	"reflect"
)

func CheckIsZero(i interface{}) bool {
	result := false

	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i).Name
		value := !(v.FieldByName(field).IsZero())
		result = value || result
	}

	return result
}
