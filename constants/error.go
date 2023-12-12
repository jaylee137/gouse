package constants

import "errors"

var (
	ErrRequiredUUID error = errors.New("id is required")
	ErrInvalidUUID  error = errors.New("uuid is invalid")

	ErrInvalidEmail error = errors.New("email is invalid")
	ErrEmailLen     error = errors.New("email must be between 8 and 32 characters long")

	ErrInvalidUsername error = errors.New("username is invalid, must be between 3 and 20 characters long and only contains letters, numbers and underscore")
	ErrInvalidPassword error = errors.New("password is invalid, must be between 8 and 32 characters long and contains at least one uppercase letter, one lowercase letter, one number and one special character")
)
