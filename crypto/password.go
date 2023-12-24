package crypto

import (
	"github.com/thuongtruong1009/gouse/constants"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", constants.ErrHashPassword
	}

	return string(hashedPassword), nil
}

func DecryptPassword(password, receivedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(receivedPassword)); err != nil {
		return constants.ErrComparePassword
	}

	return nil
}
