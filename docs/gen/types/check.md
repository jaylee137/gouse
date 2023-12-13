# Check


### IsInt
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "int")
}
```

### IsUnInt
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsUnInt(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uint")
}
```

### IsFloat
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsFloat(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "float")
}
```

### IsComplex
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsComplex(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "complex")
}
```

### IsNumber
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsNumber(v interface{}) bool {
	return IsInt(v) || IsFloat(v)
}
```

### IsString
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsString(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "string")
}
```

### IsRune
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsRune(v interface{}) bool {
	_, ok := v.(rune)
	return ok
}
```

### IsByte
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsByte(v interface{}) bool {
	_, ok := v.(byte)
	return ok
}
```

### IsUintptr
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsUintptr(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "uintptr")
}
```

### IsError
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsError(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "error")
}
```

### IsInterface
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsInterface(v interface{}) bool {
	_, ok := v.(interface{})
	return ok
}
```

### IsChannel
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsChannel(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "chan")
}
```

### IsUnsafePointer
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsUnsafePointer(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Uintptr
}
```

### IsPointer
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsPointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "*")
}
```

### IsBool
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsBool(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "bool")
}
```

### IsSlice
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsSlice(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[]")
}
```

### IsMap
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsMap(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "map")
}
```

### IsStruct
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}
```

### IsArray
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsArray(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[")
}
```

### IsFunc
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsFunc(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "func")
}
```

### IsNil
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

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
import (
	"fmt"

	"reflect"

	"strings"
)
```

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
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsUndefined(v interface{}) bool {
	return v == nil || IsEmpty(v)
}
```

### IsZero
```go
import (
	"fmt"

	"reflect"

	"strings"
)
```

```go
func IsZero(v interface{}) bool {
	return v == nil || IsEmpty(v)
}
```
