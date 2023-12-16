# Ultilities


### Select
```go
import (
	"github.com/AlecAivazis/survey/v2"
)
```

```go
func Select(label string, options []string) (string, error) {
	var selected string
	prompt := &survey.Select{
		Message: label,
		Options: options,
	}
	err := survey.AskOne(prompt, &selected)
	return selected, err
}
```

### Confirm
```go
import (
	"github.com/AlecAivazis/survey/v2"
)
```

```go
func Confirm(label string) (bool, error) {
	var confirm bool
	prompt := &survey.Confirm{
		Message: label,
	}
	err := survey.AskOne(prompt, &confirm)
	return confirm, err
}
```

### Cmd
```go
import (
	"flag"

	"fmt"

	"os"

	"os/exec"

	"runtime"

	"strconv"

	"github.com/thuongtruong1009/gouse/strings"
)
```

```go
func Cmd(defaultCmmand string, windowsCmmand ...string) {
	var cmd *exec.Cmd

	splitCmd := strings.Split(defaultCmmand, " ")

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c")
		if len(windowsCmmand) > 0 {
			cmd.Args = append(cmd.Args, windowsCmmand...)
		} else {
			cmd.Args = append(cmd.Args, splitCmd...)
		}
	} else {
		cmd = exec.Command("sh", "-c")
		cmd.Args = append(cmd.Args, splitCmd...)
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
```

### Clear
```go
import (
	"flag"

	"fmt"

	"os"

	"os/exec"

	"runtime"

	"strconv"

	"github.com/thuongtruong1009/gouse/strings"
)
```

```go
func Clear() {
	Cmd("clear", "cls")
}
```

### WithColor
```go
import (
	"flag"

	"fmt"

	"os"

	"os/exec"

	"runtime"

	"strconv"

	"github.com/thuongtruong1009/gouse/strings"
)
```

```go
func WithColor[T int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](color string, chain T) {
	fmt.Printf("%s%v\n", color, chain)
}
```

### autoDetectAction
```go
import (
	"flag"

	"fmt"

	"os"

	"os/exec"

	"runtime"

	"strconv"

	"github.com/thuongtruong1009/gouse/strings"
)
```

```go
func autoDetectAction(f interface{}) interface{} {
	switch f.(type) {
	case func(string):
		return StringAction{Action: f.(func(string))}
	case func(int):
		return IntAction{Action: f.(func(int))}
	case func(bool):
		return BoolAction{Action: f.(func(bool))}
	default:
		fmt.Println("Unsupported action type")
		os.Exit(1)
		return nil
	}
}
```

### Help
```go
import (
	"flag"

	"fmt"

	"os"

	"os/exec"

	"runtime"

	"strconv"

	"github.com/thuongtruong1009/gouse/strings"
)
```

```go
func Help(name string, options []*HelpOptions) {
	for _, option := range options {
		switch autoDetectAction(option.Action).(type) {
		case StringAction:
			flag.String(option.Opt, option.Val.(string), option.Desc)
		case IntAction:
			flag.Int(option.Opt, option.Val.(int), option.Desc)
		case BoolAction:
			flag.Bool(option.Opt, option.Val.(bool), option.Desc)
		default:
			fmt.Println("Unsupported action type")
			os.Exit(1)
		}
	}

	flag.Parse()

	if flag.NArg() > 0 && (flag.Arg(0) == "-h" || flag.Arg(0) == "--help") {
		fmt.Printf("Usage: %s [options]\n", name)
		fmt.Println("Options:")
		for _, option := range options {
			fmt.Printf("  -%s\n", option.Opt)
			fmt.Printf("        %s\n", option.Desc)
		}

		os.Exit(0)
	}

	// Run the associated actions for all set flags
	for _, option := range options {
		flag.Visit(func(f *flag.Flag) {
			if f.Name == option.Opt {
				switch a := autoDetectAction(option.Action).(type) {
				case StringAction:
					a.Action(f.Value.String())
				case IntAction:
					val, _ := strconv.Atoi(f.Value.String())
					a.Action(val)
				case BoolAction:
					val, _ := strconv.ParseBool(f.Value.String())
					a.Action(val)
				default:
					fmt.Println("Unsupported action type")
					os.Exit(1)
				}
			}
		})
	}
}
```
