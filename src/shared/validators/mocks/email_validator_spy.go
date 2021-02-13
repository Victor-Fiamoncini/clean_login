package mocks

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/validators"

// EmailValidatorSpy struct
type EmailValidatorSpy struct {
	Email        string
	IsEmailValid bool
}

// NewEmailValidatorSpy func
func NewEmailValidatorSpy() validators.IEmailValidator {
	return &EmailValidatorSpy{}
}

// GetEmail EmailValidatorSpy method
func (evs *EmailValidatorSpy) GetEmail() string {
	return evs.Email
}

// SetEmail EmailValidatorSpy method
func (evs *EmailValidatorSpy) SetEmail(email string) {
	evs.Email = email
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
func (evs *EmailValidatorSpy) Run() bool {
	return evs.IsEmailValid
}
