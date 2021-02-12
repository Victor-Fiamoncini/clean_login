package authusecase

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/encrypter"
	token_generator "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/token_generator"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
)

// AuthUseCase struct
type AuthUseCase struct {
	Email                     string
	Password                  string
	AccessToken               string
	LoadUserByEmailRepository luber.ILoadUserByEmailRepository
	Encrypter                 encrypter.IEncrypter
	TokenGenerator            token_generator.ITokenGenerator
}

// NewAuthUseCase func
func NewAuthUseCase(loadUserByEmailRepository luber.ILoadUserByEmailRepository, encrypter encrypter.IEncrypter, tokenGenerator token_generator.ITokenGenerator) IAuthUseCase {
	return &AuthUseCase{
		LoadUserByEmailRepository: loadUserByEmailRepository,
		Encrypter:                 encrypter,
		TokenGenerator:            tokenGenerator,
	}
}

// GetEmail AuthUseCase method
func (auc *AuthUseCase) GetEmail() string {
	return auc.Email
}

// SetEmail AuthUseCase method
func (auc *AuthUseCase) SetEmail(email string) {
	auc.Email = email
}

// GetPassword AuthUseCase method
func (auc *AuthUseCase) GetPassword() string {
	return auc.Password
}

// SetPassword AuthUseCase method
func (auc *AuthUseCase) SetPassword(password string) {
	auc.Password = password
}

// GetAccessToken AuthUseCase method
func (auc *AuthUseCase) GetAccessToken() string {
	return auc.AccessToken
}

// SetAccessToken AuthUseCase method
func (auc *AuthUseCase) SetAccessToken(accessToken string) {
	auc.AccessToken = accessToken
}

// Auth AuthUseCase method
func (auc *AuthUseCase) Auth(email string, password string) string {
	user := auc.LoadUserByEmailRepository.Load(email)

	if user == nil && !auc.Encrypter.Compare(password, user.GetPassword()) {
		return ""
	}

	return auc.TokenGenerator.Generate(user.GetID())
}
