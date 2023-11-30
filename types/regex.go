package types

import (
	"regexp"
	// "github.com/thuongtruong1009/gouse/shared"
)

func MatchRegex(regex, chain string) bool {
	return regexp.MustCompile(regex).MatchString(chain)
}

// func IsEmail(email string) bool {
// 	return MatchRegex(shared.EmailReg, email)
// }

// func IsUsername(username string) bool {
// 	return MatchRegex(shared.UsernameReg, username)
// }

// func IsPassword(password string) bool {
// 	return MatchRegex(shared.PasswordReg, password)
// }

// func IsPhone(phone string) bool {
// 	return MatchRegex(shared.PhoneReg, phone)
// }

// func IsUUID(uuid string) (bool, error) {
// 	return MatchRegex(shared.UUIDReg, uuid), nil
// }

// func IsEmpty(chain string) bool {
// 	return MatchRegex(EmptyReg, chain)
// }

// func IsSpace(chain string) bool {
// 	return MatchRegex(SpaceReg, chain)
// }