package crypto

import "github.com/thuongtruong1009/gouse/crypto"

func cryptoEncryptFile() {
	crypto.EncryptFile("sample.txt", []byte("password"))
	println("File content encrypted")
}

func cryptoDecryptFile() {
	crypto.DecryptFile("sample.txt", []byte("password"))
	println("File content decrypted")
}
