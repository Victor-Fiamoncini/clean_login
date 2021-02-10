package validators

// EmailValidatorSpy struct
type EmailValidatorSpy struct {
	IsEmailValid bool
}

// NewEmailValidatorSpy func
func NewEmailValidatorSpy() IEmailValidator {
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
	// var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	return evs.IsEmailValid
}
