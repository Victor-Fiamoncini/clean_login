package errors

import "errors"

// DefaultError struct
type DefaultError struct {
	Error error
	Name  string
}

// NewDefaultError func
func NewDefaultError(reason string) IDefaultError {
	return &DefaultError{
		Error: errors.New("Error with: " + reason),
		Name:  "DefaultError",
	}
}

// GetError DefaultError method
func (de *DefaultError) GetError() error {
	return de.Error
}

// SetError DefaultError method
func (de *DefaultError) SetError(err error) {
	de.Error = err
}

// GetName DefaultError method
func (de *DefaultError) GetName() string {
	return de.Name
}

// SetName DefaultError method
func (de *DefaultError) SetName(name string) {
	de.Name = name
}
