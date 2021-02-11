package authusecase

// AuthUseCaseSpy struct
type AuthUseCaseSpy struct {
	Email       string
	Password    string
	AccessToken string
}

// NewAuthUseCaseSpy func
func NewAuthUseCaseSpy() IAuthUseCase {
	return &AuthUseCaseSpy{}
}

// GetEmail AuthUseCaseSpy method
func (aucs *AuthUseCaseSpy) GetEmail() string {
	return aucs.Email
}

// SetEmail AuthUseCaseSpy method
func (aucs *AuthUseCaseSpy) SetEmail(email string) {
	aucs.Email = email
}

// GetPassword AuthUseCaseSpy method
func (aucs *AuthUseCaseSpy) GetPassword() string {
	return aucs.Password
}

// SetPassword AuthUseCaseSpy method
func (aucs *AuthUseCaseSpy) SetPassword(password string) {
	aucs.Password = password
}

// GetAccessToken AuthUseCaseSpy method
func (aucs *AuthUseCaseSpy) GetAccessToken() string {
	return aucs.AccessToken
}

// SetAccessToken AuthUseCaseSpy method
func (aucs *AuthUseCaseSpy) SetAccessToken(accessToken string) {
	aucs.AccessToken = accessToken
}

// Auth AuthUseCaseSpy method
func (aucs *AuthUseCaseSpy) Auth(email string, password string) string {
	aucs.Email = email
	aucs.Password = password

	return aucs.AccessToken
}
