package mocks

// IAuthUseCaseSpy interface
type IAuthUseCaseSpy interface {
	Auth(email string, password string) string
}

// AuthUseCaseSpy struct
type AuthUseCaseSpy struct {
	Email       string
	Password    string
	AccessToken string
}

// NewAuthUseCaseSpy func
func NewAuthUseCaseSpy() *AuthUseCaseSpy {
	return &AuthUseCaseSpy{}
}

// Auth AuthUseCaseSpy method
func (auc *AuthUseCaseSpy) Auth(email string, password string) string {
	auc.Email = email
	auc.Password = password

	return auc.AccessToken
}
