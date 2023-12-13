package constants

import "errors"

var (
	ErrRequiredUUID error = errors.New("uuid is required")
	ErrInvalidUUID  error = errors.New("uuid is invalid")
)

var (
	ErrInvalidEmail error = errors.New("email is invalid")
	ErrEmailLen     error = errors.New("email must be between 8 and 32 characters long")
)

var (
	ErrInvalidUsername error = errors.New("username is invalid, must be between 3 and 20 characters long and only contains letters, numbers and underscore")
)

var (
	ErrInvalidPassword      error = errors.New("password is invalid, must be between 8 and 32 characters long and contains at least one lowercase letter, one uppercase letter, one digit and one special character")
	ErrPasswordLen          error = errors.New("password must be between 8 and 32 characters long")
	ErrPasswordEmptyLower   error = errors.New("password must contain at least one lowercase letter")
	ErrPasswordEmptyUpper   error = errors.New("password must contain at least one uppercase letter")
	ErrPasswordEmptyDigit   error = errors.New("password must contain at least one digit")
	ErrPasswordEmptySpecial error = errors.New("password must contain at least one special character")
)

var (
	ErrInvalidPhone error = errors.New("phone is invalid")
)
