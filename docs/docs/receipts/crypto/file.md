# File

## Imports

```go
import (
	"github.com/thuongtruong1009/gouse/crypto"
)
```
## Functions


### SampleCryptoEncryptFile

```go
func SampleCryptoEncryptFile() {
	crypto.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
}
```

### SampleCryptoDecryptFile

```go
func SampleCryptoDecryptFile() {
	crypto.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
}
```
