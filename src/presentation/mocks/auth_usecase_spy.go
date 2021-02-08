package mocks

// IAuthUseCaseSpy interface
type IAuthUseCaseSpy interface {
	Auth()
}

// AuthUseCaseSpy struct
type AuthUseCaseSpy struct {
	Email    string
	Password string
}

// NewAuthUseCaseSpy func
func NewAuthUseCaseSpy() *AuthUseCaseSpy {
	return &AuthUseCaseSpy{}
}

// Auth AuthUseCaseSpy method
func (auc *AuthUseCaseSpy) Auth(email string, password string) {
	auc.Email = email
	auc.Password = password
}
