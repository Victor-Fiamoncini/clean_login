package errors

import "errors"

// InvalidParamError struct
type InvalidParamError struct {
	Error error
	Name  string
}

// NewInvalidParamError func
func NewInvalidParamError(param string) IDefaultError {
	return &InvalidParamError{
		Error: errors.New("Invalid param: " + param),
		Name:  "InvalidParamError",
	}
}

// GetError InvalidParamError method
func (ipe *InvalidParamError) GetError() error {
	return ipe.Error
}

// SetError InvalidParamError method
func (ipe *InvalidParamError) SetError(err error) {
	ipe.Error = err
}

// GetName InvalidParamError method
func (ipe *InvalidParamError) GetName() string {
	return ipe.Name
}

// SetName InvalidParamError method
func (ipe *InvalidParamError) SetName(name string) {
	ipe.Name = name
}
