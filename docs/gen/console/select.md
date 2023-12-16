# Select


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
