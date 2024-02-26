# Hash

## Imports

```go
import (
	"fmt"
	"github.com/thuongtruong1009/gouse/crypto"
)
```
## Functions


### SampleCryptoEncryptPassword

```go
func SampleCryptoEncryptPassword() {
	data := "This is a sample data"

	encryptedData, err := crypto.EncryptPassword(data)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}
	fmt.Println("Raw data:", string(data))
	fmt.Println("Encrypted data:", string(encryptedData))
}
```

### SampleCryptoDecryptPassword

```go
func SampleCryptoDecryptPassword() {
	data := "$2a$10$bcA002IOHi5SYHNH4lmIbuHjHplGl7TQZ.MznNrL1N70vAi7ovTa2"
	err := crypto.DecryptPassword(data, "This is a sample data")
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}
	println("Password matched")
}
```
