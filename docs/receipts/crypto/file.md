# File

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/crypto"
)
```
## Functions


### cryptoEncryptFile

```go
func cryptoEncryptFile() {
	crypto.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
}
```

### cryptoDecryptFile

```go
func cryptoDecryptFile() {
	crypto.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
}
```
