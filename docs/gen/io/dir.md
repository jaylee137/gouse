# Dir


### IsExistDir
```go
import (
	"os"

	"path/filepath"
)
```

```go
func IsExistDir(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return false, err
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}
```

### CreateDir
```go
import (
	"os"

	"path/filepath"
)
```

```go
func CreateDir(dir string) error {
	isExist, err := IsExistDir(dir)
	if err != nil {
		return err
	}
	if !isExist {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
```

### RemoveDir
```go
import (
	"os"

	"path/filepath"
)
```

```go
func RemoveDir(dir string) error {
	_, err := IsExistDir(dir)
	if err != nil {
		return err
	}

	err = os.RemoveAll(dir)
	if err != nil {
		return err
	} else {
		return nil
	}
}
```

### LsDir
```go
import (
	"os"

	"path/filepath"
)
```

```go
func LsDir(dir string) ([]string, error) {
	_, err := IsExistDir(dir)
	if err != nil {
		return nil, err
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}

	return result, nil
}
```

### HierarchyDir
```go
import (
	"os"

	"path/filepath"
)
```

```go
func HierarchyDir(dir string) ([]string, error) {
	_, err := IsExistDir(dir)
	if err != nil {
		return nil, err
	}

	var result []string
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		result = append(result, path)
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
```

### CurrentDir
```go
import (
	"os"

	"path/filepath"
)
```

```go
func CurrentDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}
```