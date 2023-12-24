package crypto

import (
	"encoding/base64"
)

func Encode(input []byte) ([]byte, error) {
	encodedData := base64.StdEncoding.EncodeToString(input)
	return []byte(encodedData), nil
}

func Decode(input []byte) ([]byte, error) {
	decodedData, err := base64.StdEncoding.DecodeString(string(input))
	if err != nil {
		return nil, err
	}
	return decodedData, nil
}
