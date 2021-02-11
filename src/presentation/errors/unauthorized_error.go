package errors

import (
	"errors"

	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
)

// UnauthorizedError struct
type UnauthorizedError struct {
	Error error
	Name  string
}

// NewUnauthorizedError func
func NewUnauthorizedError() shared_custom_errors.IDefaultError {
	return &UnauthorizedError{
		Error: errors.New("Unauthorized"),
		Name:  "UnauthorizedError",
	}
}

// GetError UnauthorizedError method
func (ue *UnauthorizedError) GetError() error {
	return ue.Error
}

// SetError UnauthorizedError method
func (ue *UnauthorizedError) SetError(err error) {
	ue.Error = err
}

// GetName UnauthorizedError method
func (ue *UnauthorizedError) GetName() string {
	return ue.Name
}

// SetName UnauthorizedError method
func (ue *UnauthorizedError) SetName(name string) {
	ue.Name = name
}
