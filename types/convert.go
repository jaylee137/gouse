package types

import (
	"fmt"
	"reflect"
)

func StructToString(data interface{}) string {
	v := reflect.ValueOf(data)
	t := v.Type()

	var fields []string

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		fields = append(fields, fmt.Sprintf("%s: %v", fieldName, fieldValue))
	}

	// replace with array.Join()
	var result string
	for i, v := range fields {
		result += v
		if i != len(fields)-1 {
			result += ", "
		}
	}

	return fmt.Sprintf("%s{%s}", t.Name(), result)
}
