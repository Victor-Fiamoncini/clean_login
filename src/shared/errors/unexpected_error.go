package errors

import "errors"

// UnexpectedError struct
type UnexpectedError struct {
	Error error
	Name  string
}

// NewUnexpectedError func
func NewUnexpectedError(reason string) IDefaultError {
	return &UnexpectedError{
		Error: errors.New("Unexpected Error with: " + reason),
		Name:  "UnexpectedError",
	}
}

// GetError UnexpectedError method
func (ue *UnexpectedError) GetError() error {
	return ue.Error
}

// SetError UnexpectedError method
func (ue *UnexpectedError) SetError(err error) {
	ue.Error = err
}

// GetName UnexpectedError method
func (ue *UnexpectedError) GetName() string {
	return ue.Name
}

// SetName UnexpectedError method
func (ue *UnexpectedError) SetName(name string) {
	ue.Name = name
}
