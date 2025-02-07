package storage

import "errors"

var (
	ErrUserExists   = errors.New("user alredy exists")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFound  = errors.New("app not found")
)
