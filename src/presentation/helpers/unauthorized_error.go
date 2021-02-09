package helpers

import "errors"

// UnauthorizedError struct
type UnauthorizedError struct {
	Error error
	Name  string
}

// NewUnauthorizedError func
func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{
		Error: errors.New("Unauthorized"),
		Name:  "UnauthorizedError",
	}
}
