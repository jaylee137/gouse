
### TestConsole
```go
import (
	"reflect"

	"testing"
)
```

```go
func TestConsole(t *testing.T) {
	colors := []string{DEFAULT_FG, RED_FG, GREEN_FG, YELLOW_FG, BLUE_FG, MAGENTA_FG, CYAN_FG, WHITE_FG, ORANGE_FG, PURPLE_FG, PINK_FG}

	for _, color := range colors {
		if reflect.TypeOf(color).Kind() != reflect.String {
			t.Errorf("%s is not string", color)
		} else if len(color) == 0 {
			t.Errorf("%s is empty", color)
		} else if color[0] != '\033' {
			t.Errorf("%s is not a color", color)
		} else if color[len(color)-1] != 'm' {
			t.Errorf("%s is not a color", color)
		}
	}
}
```
