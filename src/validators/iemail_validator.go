package validators

// IEmailValidator interface
type IEmailValidator interface {
	GetIsEmailValid() bool
	SetIsEmailValid(isEmailValid bool)

	Run(email string) bool
}
