package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// EmailValidator struct
type EmailValidator struct {
	Email        string
	IsEmailValid bool
}

// NewEmailValidator func
func NewEmailValidator() IEmailValidator {
	return &EmailValidator{}
}

// GetEmail EmailValidator method
func (ev *EmailValidator) GetEmail() string {
	return ev.Email
}

// SetEmail EmailValidator method
func (ev *EmailValidator) SetEmail(email string) {
	ev.Email = email
}

// GetIsEmailValid EmailValidator method
func (ev *EmailValidator) GetIsEmailValid() bool {
	return ev.IsEmailValid
}

// SetIsEmailValid EmailValidator method
func (ev *EmailValidator) SetIsEmailValid(isEmailValid bool) {
	ev.IsEmailValid = isEmailValid
}

// Run NewEmailValidator method
func (ev *EmailValidator) Run() bool {
	if ev.Email == "" {
		return false
	}

	err := validation.Validate(ev.Email, is.Email)

	ev.IsEmailValid = err == nil

	return ev.IsEmailValid
}
