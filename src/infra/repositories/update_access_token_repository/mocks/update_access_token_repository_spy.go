package mocks

import (
	uatr "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository"
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
)

// UpdateAccessTokenRepositorySpy struct
type UpdateAccessTokenRepositorySpy struct {
	UserID      string
	AccessToken string
}

// NewUpdateAccessTokenRepositorySpy func
func NewUpdateAccessTokenRepositorySpy() uatr.IUpdateAccessTokenRepository {
	return &UpdateAccessTokenRepositorySpy{}
}

// GetUserID UpdateAccessTokenRepositorySpy method
func (uatrs *UpdateAccessTokenRepositorySpy) GetUserID() string {
	return uatrs.UserID
}

// SetUserID UpdateAccessTokenRepositorySpy method
func (uatrs *UpdateAccessTokenRepositorySpy) SetUserID(userID string) {
	uatrs.UserID = userID
}

// GetAccessToken UpdateAccessTokenRepositorySpy method
func (uatrs *UpdateAccessTokenRepositorySpy) GetAccessToken() string {
	return uatrs.AccessToken
}

// SetAccessToken UpdateAccessTokenRepositorySpy method
func (uatrs *UpdateAccessTokenRepositorySpy) SetAccessToken(accessToken string) {
	uatrs.AccessToken = accessToken
}

// Update UpdateAccessTokenRepositorySpy method
func (uatrs *UpdateAccessTokenRepositorySpy) Update() shared_custom_errors.IDefaultError {
	return nil
}
