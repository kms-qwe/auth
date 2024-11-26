package model

import "errors"

var (
	// ErrorUserNotFound represents user error user not found in cache
	ErrorUserNotFound = errors.New("user not found")
)
