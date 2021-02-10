package validators

import "regexp"

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
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return emailRegex.MatchString(email)
}
