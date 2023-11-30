package constants

import "errors"

const (
	KeyStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	ErrRequiredUUID error = errors.New("id is required")
	ErrInvalidUUID  error = errors.New("uuid is invalid")
)
