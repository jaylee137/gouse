package structs

import (
	"reflect"
)

func Clone(structInstance interface{}) interface{} {
	structValue := reflect.ValueOf(structInstance)

	// Create a new struct instance
	newStructValue := reflect.New(structValue.Type()).Elem()

	// Copy values from the original struct to the new struct
	for i := 0; i < structValue.NumField(); i++ {
		fieldValue := structValue.Field(i)
		newStructValue.Field(i).Set(fieldValue)
	}

	return newStructValue.Interface()
}

func Has(structInstance interface{}, fieldName string) bool {
	structType := reflect.TypeOf(structInstance)
	_, ok := structType.FieldByName(fieldName)
	return ok
}

func HasEmpty(structInstance interface{}, fieldName string) bool {
	structValue := reflect.ValueOf(structInstance)
	fieldValue := structValue.FieldByName(fieldName)

	return reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface())
}
