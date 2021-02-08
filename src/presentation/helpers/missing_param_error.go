package helpers

import "errors"

// MissingParamError struct
type MissingParamError struct {
	Error error
	Name  string
}

// NewMissingParamError func
func NewMissingParamError(param string) *MissingParamError {
	return &MissingParamError{
		Error: errors.New("Missing param: " + param),
		Name:  "MissingParamError",
	}
}
