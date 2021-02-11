package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// EmailValidator struct
type EmailValidator struct {
	IsEmailValid bool
}

// NewEmailValidator func
func NewEmailValidator() IEmailValidator {
	return &EmailValidator{}
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
func (ev *EmailValidator) Run(email string) bool {
	err := validation.Validate(email, is.Email)

	return err == nil
}
