package tokengenerator

import (
	"time"

	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	jwt "github.com/dgrijalva/jwt-go"
)

// TokenGenerator struct
type TokenGenerator struct {
	UserID      string
	AccessToken string
	Secret      string
}

// NewTokenGenerator func
func NewTokenGenerator(secret string) ITokenGenerator {
	return &TokenGenerator{
		Secret: secret,
	}
}

// GetUserID TokenGenerator method
func (tg *TokenGenerator) GetUserID() string {
	return tg.UserID
}

// SetUserID TokenGenerator method
func (tg *TokenGenerator) SetUserID(userID string) {
	tg.UserID = userID
}

// GetAccessToken TokenGenerator method
func (tg *TokenGenerator) GetAccessToken() string {
	return tg.AccessToken
}

// SetAccessToken TokenGenerator method
func (tg *TokenGenerator) SetAccessToken(accessToken string) {
	tg.AccessToken = accessToken
}

// GetSecret TokenGenerator method
func (tg *TokenGenerator) GetSecret() string {
	return tg.Secret
}

// SetSecret TokenGenerator method
func (tg *TokenGenerator) SetSecret(secret string) {
	tg.Secret = secret
}

// Generate TokenGenerator method
func (tg *TokenGenerator) Generate() (string, shared_custom_errors.IDefaultError) {
	if tg.Secret == "" {
		err := shared_custom_errors.NewMissingParamError("Secret")

		return "", err
	}

	if tg.UserID == "" {
		err := shared_custom_errors.NewMissingParamError("UserID")

		return "", err
	}

	claims := jwt.MapClaims{}

	claims["userId"] = tg.UserID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedKey := []byte(tg.Secret)

	generatedToken, jwtErr := tokenWithClaims.SignedString(signedKey)

	if jwtErr != nil {
		return "", shared_custom_errors.NewUnexpectedError("TokenGenerator.Generate()")
	}

	tg.AccessToken = generatedToken

	return tg.AccessToken, nil
}
