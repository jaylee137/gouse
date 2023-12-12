package shared

// const (
// 	EmptyReg   = `^$`
// 	SpaceReg   = `\s+`
// 	WordNumReg = `^[a-zA-Z0-9]+$`
// )

const (
	UsernameReg = `^[a-zA-Z0-9_]{3,20}$`
	PasswordReg = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*]).{8,32}$`
)

const (
	// PasswordLenReg = `^.{8,32}$`
	EmailLenReg = `^.{8,32}$`
)
