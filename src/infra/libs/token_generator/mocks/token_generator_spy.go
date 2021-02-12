package mocks

import (
	token_generator "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/token_generator"
)

// TokenGeneratorSpy struct
type TokenGeneratorSpy struct {
	UserID string
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
