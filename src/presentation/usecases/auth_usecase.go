package usecases

// IAuthUseCase interface
type IAuthUseCase interface {
	Auth()
}

// AuthUseCase struct
type AuthUseCase struct {
	Email    string
	Password string
}

// NewAuthUseCase func
func NewAuthUseCase() *AuthUseCase {
	return &AuthUseCase{}
}

// Auth AuthUseCase method
func (auc *AuthUseCase) Auth(email string, password string) {
	auc.Email = email
	auc.Password = password
}
