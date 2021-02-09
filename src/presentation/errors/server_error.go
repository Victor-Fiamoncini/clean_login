package errors

import "errors"

// ServerError struct
type ServerError struct {
	Error error
	Name  string
}

// NewServerError func
func NewServerError() IDefaultError {
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
