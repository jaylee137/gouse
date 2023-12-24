package types

import (
	"fmt"
	"reflect"
	"strings"
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

func IsChannel(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "chan")
}

func IsUnsafePointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Uintptr
}

func IsPointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "*")
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
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

func IsArray(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[")
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

	switch value := v.(type) {
	case string:
		return value == ""
	case int, int8, int16, int32, int64:
		return value == 0
	case uint, uint8, uint16, uint32, uint64:
		return value == 0
	case float32, float64:
		return value == 0
	case bool:
		return !value
	case []interface{}:
		return len(value) == 0
	case map[string]interface{}:
		return len(value) == 0
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
