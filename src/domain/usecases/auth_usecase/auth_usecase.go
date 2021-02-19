package authusecase

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	uatr "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository"
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/encrypter"
	token_generator "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/token_generator"
)

// AuthUseCase struct
type AuthUseCase struct {
	Email                       string
	Password                    string
	AccessToken                 string
	LoadUserByEmailRepository   luber.ILoadUserByEmailRepository
	Encrypter                   encrypter.IEncrypter
	TokenGenerator              token_generator.ITokenGenerator
	UpdateAccessTokenRepository uatr.IUpdateAccessTokenRepository
}

// NewAuthUseCase func
func NewAuthUseCase(loadUserByEmailRepository luber.ILoadUserByEmailRepository, encrypter encrypter.IEncrypter, tokenGenerator token_generator.ITokenGenerator, updateAccessTokenRepository uatr.IUpdateAccessTokenRepository) IAuthUseCase {
	return &AuthUseCase{
		LoadUserByEmailRepository:   loadUserByEmailRepository,
		Encrypter:                   encrypter,
		TokenGenerator:              tokenGenerator,
		UpdateAccessTokenRepository: updateAccessTokenRepository,
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
func (auc *AuthUseCase) Auth() (string, shared_custom_errors.IDefaultError) {
	var err shared_custom_errors.IDefaultError
	var user entities.IUser

	auc.LoadUserByEmailRepository.SetEmail(auc.GetEmail())

	user, err = auc.LoadUserByEmailRepository.Load()

	if err != nil {
		return "", err
	}

	var isPasswordValid bool

	auc.Encrypter.SetPassword(auc.GetPassword())
	auc.Encrypter.SetHashedPassword(user.GetPassword())

	isPasswordValid, err = auc.Encrypter.Compare()

	if err != nil || !isPasswordValid {
		return "", err
	}

	var accessToken string

	auc.TokenGenerator.SetUserID(user.GetID())

	accessToken, err = auc.TokenGenerator.Generate()

	if err != nil {
		return "", err
	}

	auc.UpdateAccessTokenRepository.SetUserID(user.GetID())
	auc.UpdateAccessTokenRepository.SetAccessToken(accessToken)

	err = auc.UpdateAccessTokenRepository.Update()

	if err != nil {
		return "", err
	}

	return accessToken, nil
}
