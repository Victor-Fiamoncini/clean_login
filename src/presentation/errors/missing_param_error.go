package errors

import "errors"

// MissingParamError struct
type MissingParamError struct {
	Error error
	Name  string
}

// NewMissingParamError func
func NewMissingParamError(param string) IDefaultError {
	return &MissingParamError{
		Error: errors.New("Missing param: " + param),
		Name:  "MissingParamError",
	}
}

// GetError MissingParamError method
func (mpe *MissingParamError) GetError() error {
	return mpe.Error
}

// SetError MissingParamError method
func (mpe *MissingParamError) SetError(err error) {
	mpe.Error = err
}

// GetName MissingParamError method
func (mpe *MissingParamError) GetName() string {
	return mpe.Name
}

// SetName MissingParamError method
func (mpe *MissingParamError) SetName(name string) {
	mpe.Name = name
}
