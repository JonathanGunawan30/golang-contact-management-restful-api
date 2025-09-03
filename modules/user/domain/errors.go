package domain

import "errors"

var (
	ErrUsernameAlreadyExists = errors.New("username already exists")
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidCredentials    = errors.New("username or password is wrong")
)
