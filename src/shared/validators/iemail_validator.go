package validators

// IEmailValidator interface
type IEmailValidator interface {
	GetEmail() string
	SetEmail(email string)
	GetIsEmailValid() bool
	SetIsEmailValid(isEmailValid bool)

	Run() bool
}
