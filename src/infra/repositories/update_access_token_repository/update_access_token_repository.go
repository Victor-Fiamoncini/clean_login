package updateaccesstokenrepository

import (
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
)

// UpdateAccessTokenRepository struct
type UpdateAccessTokenRepository struct {
	UserID      string
	AccessToken string
}

// NewUpdateAccessTokenRepository func
func NewUpdateAccessTokenRepository() IUpdateAccessTokenRepository {
	return &UpdateAccessTokenRepository{}
}

// GetUserID UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) GetUserID() string {
	return uatr.UserID
}

// SetUserID UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) SetUserID(userID string) {
	uatr.UserID = userID
}

// GetAccessToken UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) GetAccessToken() string {
	return uatr.AccessToken
}

// SetAccessToken UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) SetAccessToken(accessToken string) {
	uatr.AccessToken = accessToken
}

// Update UpdateAccessTokenRepository method
func (uatr *UpdateAccessTokenRepository) Update() shared_custom_errors.IDefaultError {
	return nil
}
