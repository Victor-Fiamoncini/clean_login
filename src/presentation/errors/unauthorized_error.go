package errors

import "errors"

// UnauthorizedError struct
type UnauthorizedError struct {
	Error error
	Name  string
}

// NewUnauthorizedError func
func NewUnauthorizedError() IDefaultError {
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
