package tokengenerator

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// TokenGenerator struct
type TokenGenerator struct {
	UserID      string
	AccessToken string
}

// NewTokenGenerator func
func NewTokenGenerator() ITokenGenerator {
	return &TokenGenerator{}
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

// Generate TokenGenerator method
func (tg *TokenGenerator) Generate() string {
	claims := jwt.MapClaims{}

	claims["userId"] = tg.UserID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	generatedToken, err := tokenWithClaims.SignedString("MY_JWT_SECRET")

	if err != nil {
		return ""
	}

	tg.AccessToken = generatedToken

	return tg.AccessToken
}
