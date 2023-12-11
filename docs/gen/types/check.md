# Check


### IsInt
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsInterface(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "interface")
}
```

### IsChannel
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsUnsafePointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "unsafe.Pointer")
}
```

### IsBool
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsStruct(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "struct")
}
```

### IsArray
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsArray(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "[")
}
```

### IsPointer
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsPointer(v interface{}) bool {
	return strings.Contains(fmt.Sprintf("%T", v), "*")
}
```

### IsFunc
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
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

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
func IsZero(v interface{}) bool {
	return v == nil || IsEmpty(v)
}
```

### IsUUID
```go
import (
	"fmt"

	"strings"

	"github.com/google/uuid"

	"github.com/thuongtruong1009/gouse/constants"
)
```

```go
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
```
