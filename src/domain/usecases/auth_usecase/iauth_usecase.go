package authusecase

import shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"

// IAuthUseCase interface
type IAuthUseCase interface {
	GetEmail() string
	SetEmail(email string)
	GetPassword() string
	SetPassword(password string)
	GetAccessToken() string
	SetAccessToken(accessToken string)

	Auth() (string, shared_custom_errors.IDefaultError)
}
