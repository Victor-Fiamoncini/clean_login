package mocks

import (
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	token_generator "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/token_generator"
)

// TokenGeneratorSpy struct
type TokenGeneratorSpy struct {
	UserID      string
	AccessToken string
	Secret      string
}

// NewTokenGeneratorSpy func
func NewTokenGeneratorSpy() token_generator.ITokenGenerator {
	return &TokenGeneratorSpy{}
}

// GetUserID TokenGeneratorSpy method
func (tgs *TokenGeneratorSpy) GetUserID() string {
	return tgs.UserID
}

// SetUserID TokenGeneratorSpy method
func (tgs *TokenGeneratorSpy) SetUserID(userID string) {
	tgs.UserID = userID
}

// GetAccessToken TokenGeneratorSpy method
func (tgs *TokenGeneratorSpy) GetAccessToken() string {
	return tgs.AccessToken
}

// SetAccessToken TokenGeneratorSpy method
func (tgs *TokenGeneratorSpy) SetAccessToken(accessToken string) {
	tgs.AccessToken = accessToken
}

// GetSecret TokenGeneratorSpy method
func (tgs *TokenGeneratorSpy) GetSecret() string {
	return tgs.Secret
}

// SetSecret TokenGeneratorSpy method
func (tgs *TokenGeneratorSpy) SetSecret(secret string) {
	tgs.Secret = secret
}

// Generate TokenGeneratorSpy method
func (tgs *TokenGeneratorSpy) Generate() (string, shared_custom_errors.IDefaultError) {
	return tgs.AccessToken, nil
}
