package shared

// const (
// 	EmptyReg   = `^$`
// 	SpaceReg   = `\s+`
// 	WordNumReg = `^[a-zA-Z0-9]+$`
// )

const (
	UsernameReg = `^[a-zA-Z0-9_]{3,20}$`
)

const (
	PasswordLenReg     = `^.{8,32}$`
	PasswordLowerReg   = `[a-z]`
	PasswordUpperReg   = `[A-Z]`
	PasswordDigitReg   = `\d`
	PasswordSpecialReg = `[!@#$%^&*]`
)

const (
	EmailLenReg = `^.{8,32}$`
)

const (
	PhoneReg = `^\+\d{1,2}\s?\(\d{1,4}\)\s?\d{1,6}-\d{1,6}$`
)
