package domain

import "errors"

var (
	ErrNotFoundUser = errors.New("not found user")
	ErrUnexpected   = errors.New("unexpected")
)
