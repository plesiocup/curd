package util

import "reflect"

func HasEmptyField(obj interface{}, fieldNames ...string) bool {
	value := reflect.ValueOf(obj).Elem()
	for _, fieldName := range fieldNames {
		fieldValue := value.FieldByName(fieldName).String()
		if fieldValue == "" {
			return true
		}
	}
	return false
}
