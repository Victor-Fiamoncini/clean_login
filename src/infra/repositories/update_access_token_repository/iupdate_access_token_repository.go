package updateaccesstokenrepository

import shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"

// IUpdateAccessTokenRepository interface
type IUpdateAccessTokenRepository interface {
	GetUserID() string
	SetUserID(userID string)
	GetAccessToken() string
	SetAccessToken(accessToken string)

	Update() shared_custom_errors.IDefaultError
}
