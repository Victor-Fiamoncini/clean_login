package usecases

// IAuthUseCase interface
type IAuthUseCase interface {
	Auth(email string, password string) string
}

// AuthUseCase struct
type AuthUseCase struct {
	Email       string
	Password    string
	AccessToken string
}

// NewAuthUseCase func
func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{}
}

// Auth AuthUseCase method
func (auc *AuthUseCase) Auth(email string, password string) string {
	auc.Email = email
	auc.Password = password

	return auc.AccessToken
}
