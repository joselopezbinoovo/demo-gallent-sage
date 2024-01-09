package utils

import (
	"reflect"
)

func GetType(element interface{}) reflect.Type {
	reflectType := reflect.ValueOf(element).Type()
	for reflectType.Kind() == reflect.Slice || reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return reflectType
}
