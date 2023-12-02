package shared

const (
	EmptyReg   = `^$`
	SpaceReg   = `\s+`
	WordNumReg = `^[a-zA-Z0-9]+$`
)

const (
	UsernameLenReg = `^.{4,20}$`
	PasswordLenReg = `^.{8,32}$`
	EmailLenReg    = `^.{8,32}$`
)
