package roleman

import "errors"

var (
	ErrNotFound  = errors.New("role.not_found")
	UserNotFound = errors.New("user.role.not_found")
)
