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

func StructToMap(data interface{}) map[string]interface{} {
	v := reflect.ValueOf(data)
	t := v.Type()

	result := make(map[string]interface{})

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

func StringToInt(data string) int {
	var result int
	fmt.Sscanf(data, "%d", &result)
	return result
}

func StringToFloat(data string) float64 {
	var result float64
	fmt.Sscanf(data, "%f", &result)
	return result
}

func StringToBool(data string) bool {
	var result bool
	fmt.Sscanf(data, "%t", &result)
	return result
}

func IntToString(data int) string {
	return fmt.Sprintf("%d", data)
}

func FloatToString(data float64) string {
	return fmt.Sprintf("%f", data)
}

func BoolToString(data bool) string {
	return fmt.Sprintf("%t", data)
}

func InterfaceToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}

// this four func not have example code
func BytesToString(data []byte) string {
	return string(data)
}

func StringToBytes(data string) []byte {
	return []byte(data)
}

func StringsToBytes(data []string) []byte {
	var result []byte
	for _, v := range data {
		result = append(result, StringToBytes(v)...)
	}
	return result
}

func RuneToString(data rune) string {
	return string(data)
}
