# Crypto

## Imports

```go
import (
	"fmt"	"github.com/thuongtruong1009/gouse/crypto")
```
## Functions


### cryptoEncode

```go
func cryptoEncode() {
	data := []byte("This is a sample data")

	encodedData, err := crypto.Encode(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Println("Raw data:", string(data))
	fmt.Println("Encoded data:", string(encodedData))
}```

### cryptoDecode

```go
func cryptoDecode() {
	data := []byte("VGhpcyBpcyBhIHNhbXBsZSBkYXRh")
	decodedData, err := crypto.Decode(data)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}
	fmt.Println("Raw encoded data:", string(data))
	fmt.Println("Decoded data:", string(decodedData))
}```

### cryptoEncryptPassword

```go
func cryptoEncryptPassword() {
	data := "This is a sample data"

	encryptedData, err := crypto.EncryptPassword(data)
	if err != nil {
		fmt.Println("Error encrypting data:", err)
		return
	}
	fmt.Println("Raw data:", string(data))
	fmt.Println("Encrypted data:", string(encryptedData))
}```

### cryptoDecryptPassword

```go
func cryptoDecryptPassword() {
	data := "$2a$10$bcA002IOHi5SYHNH4lmIbuHjHplGl7TQZ.MznNrL1N70vAi7ovTa2"
	err := crypto.DecryptPassword(data, "This is a sample data")
	if err != nil {
		fmt.Println("Error decrypting data:", err)
		return
	}
	println("Password matched")
}```

### cryptoEncryptFile

```go
func cryptoEncryptFile() {
	crypto.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
}```

### cryptoDecryptFile

```go
func cryptoDecryptFile() {
	crypto.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
}```
