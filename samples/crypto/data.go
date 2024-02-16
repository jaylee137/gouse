package crypto

import (
	"fmt"

	"github.com/thuongtruong1009/gouse/crypto"
)

func cryptoEncode() {
	data := []byte("This is a sample data")

	encodedData, err := crypto.Encode(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	fmt.Println("Raw data:", string(data))
	fmt.Println("Encoded data:", string(encodedData))
}

func cryptoDecode() {
	data := []byte("VGhpcyBpcyBhIHNhbXBsZSBkYXRh")
	decodedData, err := crypto.Decode(data)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return
	}
	fmt.Println("Raw encoded data:", string(data))
	fmt.Println("Decoded data:", string(decodedData))
}
