# Convert

## Imports

```go
import (
	"fmt"

	"reflect"

	"strings"
)
```
## Functions


### IsInt

```go
func IsInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "int")
}
```

### IsUnInt

```go
func IsUnInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uint")
}
```

### IsFloat

```go
func IsFloat(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "float")
}
```

### IsComplex

```go
func IsComplex(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "complex")
}
```

### IsNumber

```go
func IsNumber(v interface{}) bool {
	return IsInt(v) || IsFloat(v)
}
```

### IsString

```go
func IsString(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "string")
}
```

### IsRune

```go
func IsRune(v interface{}) bool {
	_, ok := v.(rune)
	return ok
}
```

### IsByte

```go
func IsByte(v interface{}) bool {
	_, ok := v.(byte)
	return ok
}
```

### IsUintptr

```go
func IsUintptr(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uintptr")
}
```

### IsError

```go
func IsError(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "error")
}
```

### IsInterface

```go
func IsInterface(v interface{}) bool {
	_, ok := v.(interface{})
	return ok
}
```

### IsChannel

```go
func IsChannel(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "chan")
}
```

### IsUnsafePointer

```go
func IsUnsafePointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Uintptr
}
```

### IsPointer

```go
func IsPointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "*")
}
```

### IsBool

```go
func IsBool(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "bool")
}
```

### IsSlice

```go
func IsSlice(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[]")
}
```

### IsMap

```go
func IsMap(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "map")
}
```

### IsStruct

```go
func IsStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}
```

### IsArray

```go
func IsArray(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[")
}
```

### IsFunc

```go
func IsFunc(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "func")
}
```

### IsNil

```go
func IsNil(v interface{}) bool {
	return v == nil

	// 	v := reflect.ValueOf(value)
	// 	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
	// 		return v.IsNil()
	// 	}

	// 	return false
}
```

### IsEmpty

```go
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
```

### IsUndefined

```go
func IsUndefined(v interface{}) bool {
	return v == nil || IsEmpty(v)
}
```

### IsZero

```go
func IsZero(v interface{}) bool {
	return v == nil || IsEmpty(v)
}
```

### StructToString

```go
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
```

### StructToMap

```go
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
```

### StringToInt

```go
func StringToInt(data string) int {
	var result int
	fmt.Sscanf(data, "%d", &result)
	return result
}
```

### StringToFloat

```go
func StringToFloat(data string) float64 {
	var result float64
	fmt.Sscanf(data, "%f", &result)
	return result
}
```

### StringToBool

```go
func StringToBool(data string) bool {
	var result bool
	fmt.Sscanf(data, "%t", &result)
	return result
}
```

### StringToBytes

```go
func StringToBytes(data string) []byte {
	return []byte(data)
}
```

### StringsToBytes

```go
func StringsToBytes(data []string) []byte {
	var result []byte
	for _, v := range data {
		result = append(result, StringToBytes(v)...)
	}
	return result
}
```

### IntToString

```go
func IntToString(data int) string {
	return fmt.Sprintf("%d", data)
}
```

### FloatToString

```go
func FloatToString(data float64) string {
	return fmt.Sprintf("%f", data)
}
```

### BoolToString

```go
func BoolToString(data bool) string {
	return fmt.Sprintf("%t", data)
}
```

### InterfaceToString

```go
func InterfaceToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}
```

### BytesToString

```go
func BytesToString(data []byte) string {
	return string(data)
}
```

### RuneToString

```go
func RuneToString(data rune) string {
	return string(data)
}
```

### MapAsString

```go
func MapAsString[T string | []string](data map[string]T) string {
	var result string

	for key, value := range data {
		result += fmt.Sprintf("%s: %s\n", key, value)
	}

	return result
}
```
