package mocks

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/validators"

// EmailValidatorSpy struct
type EmailValidatorSpy struct {
	IsEmailValid bool
}

// NewEmailValidatorSpy func
func NewEmailValidatorSpy() validators.IEmailValidator {
	return &EmailValidatorSpy{}
}

// GetIsEmailValid EmailValidatorSpy method
func (evs *EmailValidatorSpy) GetIsEmailValid() bool {
	return evs.IsEmailValid
}

// SetIsEmailValid EmailValidatorSpy method
func (evs *EmailValidatorSpy) SetIsEmailValid(isEmailValid bool) {
	evs.IsEmailValid = isEmailValid
}

// Run NewEmailValidatorSpy method
func (evs *EmailValidatorSpy) Run(email string) bool {
	return evs.IsEmailValid
}
