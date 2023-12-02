package constants

import "errors"

var (
	ErrRequiredUUID error = errors.New("id is required")
	ErrInvalidUUID  error = errors.New("uuid is invalid")
)
