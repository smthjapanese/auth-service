package entity

import "errors"

var (
	ErrInvalidUsername      = errors.New("invalid username")
	ErrInvalidPass          = errors.New("invalid password")
	ErrUsernameAlreadyExist = errors.New("username already exists")
	ErrUsernameNotFound     = errors.New("username not found")
)
