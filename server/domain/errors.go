package domain

import "errors"

var (
	ErrNotFoundUser    = errors.New("not found user")
	ErrNotFoundProblem = errors.New("not found problem")
	ErrUnexpected      = errors.New("unexpected")
)
