package tokengenerator

import shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"

// ITokenGenerator interface
type ITokenGenerator interface {
	GetUserID() string
	SetUserID(userID string)
	GetAccessToken() string
	SetAccessToken(accessToken string)
	GetSecret() string
	SetSecret(secret string)

	Generate() (string, shared_custom_errors.IDefaultError)
}
