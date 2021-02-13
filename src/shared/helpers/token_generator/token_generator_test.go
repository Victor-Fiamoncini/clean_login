package tokengenerator_test

import (
	"testing"

	token_generator "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/token_generator"
	"github.com/go-playground/assert/v2"
)

func makeSut() token_generator.ITokenGenerator {
	tokenGenerator := token_generator.NewTokenGenerator("MY_SECRET")

	tokenGenerator.SetUserID("any_id")

	return tokenGenerator
}

func TestShouldReturnNullIfJWTReturnsNull(t *testing.T) {
	sut := makeSut()

	token := sut.Generate()

	assert.Equal(t, "", token)
}

func TestShouldReturnANewTokenIfJWTReturnsANewToken(t *testing.T) {
	sut := makeSut()

	token := sut.Generate()

	assert.Equal(t, sut.GetAccessToken(), token)
}

func TestShouldReturnNullIfRequiredParamsAreNotProvided(t *testing.T) {
	sut := token_generator.NewTokenGenerator("")

	token := sut.Generate()

	assert.Equal(t, "", token)
}
