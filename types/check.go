package types

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/thuongtruong1009/gouse/constants"
)

func IsInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "int")
}

func IsUnInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uint")
}

func IsFloat(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "float")
}

func IsComplex(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "complex")
}

func IsNumber(v interface{}) bool {
	return IsInt(v) || IsFloat(v)
}

func IsString(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "string")
}

func IsRune(v interface{}) bool {
	_, ok := v.(rune)
	return ok
}

func IsByte(v interface{}) bool {
	_, ok := v.(byte)
	return ok
}

func IsUintptr(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uintptr")
}

func IsError(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "error")
}

func IsInterface(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "interface")
}

func IsChannel(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "chan")
}

func IsUnsafePointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "unsafe.Pointer")
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

	// 	v := reflect.ValueOf(value)
	// 	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
	// 		return v.IsNil()
	// 	}

	// 	return false
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

func IsUndefined(v interface{}) bool {
	return v == nil || IsEmpty(v)
}

func IsZero(v interface{}) bool {
	return v == nil || IsEmpty(v)
}

func IsUUID(input string) (bool, error) {
	if input == "" {
		return false, constants.ErrRequiredUUID
	}

	_, err := uuid.Parse(input)
	if err != nil {
		return false, constants.ErrInvalidUUID
	}

	return true, nil
}
