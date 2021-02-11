package errors

import (
	"errors"

	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
)

// ServerError struct
type ServerError struct {
	Error error
	Name  string
}

// NewServerError func
func NewServerError() shared_custom_errors.IDefaultError {
	return &ServerError{
		Error: errors.New("Internal Error"),
		Name:  "ServerError",
	}
}

// GetError ServerError method
func (se *ServerError) GetError() error {
	return se.Error
}

// SetError ServerError method
func (se *ServerError) SetError(err error) {
	se.Error = err
}

// GetName ServerError method
func (se *ServerError) GetName() string {
	return se.Name
}

// SetName ServerError method
func (se *ServerError) SetName(name string) {
	se.Name = name
}
