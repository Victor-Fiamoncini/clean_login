package mocks

import authusecase "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases/auth_usecase"

// AuthUseCaseSpy struct
type AuthUseCaseSpy struct {
	Email       string
	Password    string
	AccessToken string
}

// NewAuthUseCaseSpy func
func NewAuthUseCaseSpy() authusecase.IAuthUseCase {
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
func (aucs *AuthUseCaseSpy) Auth() string {
	return aucs.AccessToken
}
