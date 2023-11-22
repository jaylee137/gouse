package types

import (
	"fmt"
	"strings"
)

func IsInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "int")
}

func IsFloat(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "float")
}

func IsString(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "string")
}

func IsBool(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "bool")
}

func IsSlice(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[]")
}

func IsMap(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "map")
}

func IsStruct(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "struct")
}

func IsArray(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[")
}

func IsPointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "*")
}

func IsFunc(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "func")
}

func IsNil(v interface{}) bool {
	return v == nil
}

func IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}
	switch v.(type) {
	case string:
		return v.(string) == ""
	case int:
		return v.(int) == 0
	case int8:
		return v.(int8) == 0
	case int16:
		return v.(int16) == 0
	case int32:
		return v.(int32) == 0
	case int64:
		return v.(int64) == 0
	case uint:
		return v.(uint) == 0
	case uint8:
		return v.(uint8) == 0
	case uint16:
		return v.(uint16) == 0
	case uint32:
		return v.(uint32) == 0
	case uint64:
		return v.(uint64) == 0
	case float32:
		return v.(float32) == 0
	case float64:
		return v.(float64) == 0
	case bool:
		return !v.(bool)
	case []interface{}:
		return len(v.([]interface{})) == 0
	case map[string]interface{}:
		return len(v.(map[string]interface{})) == 0
	default:
		return false
	}
}