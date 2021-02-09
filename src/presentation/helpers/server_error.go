package helpers

import "errors"

// ServerError struct
type ServerError struct {
	Error error
	Name  string
}

// NewServerError func
func NewServerError() *ServerError {
	return &ServerError{
		Error: errors.New("Internal Error"),
		Name:  "ServerError",
	}
}
