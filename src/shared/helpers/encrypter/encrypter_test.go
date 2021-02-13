package encrypter

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func makeSut() (IEncrypter, string) {
	const validPasswordHash = "$2a$09$ekW4Ggqbaixn3x/qyzcreex5RoFIaPAMPLrDaxDZg0QE0Qe008nOW"

	return NewEncrypter(), validPasswordHash
}

func TestShouldReturnTrueIfBCryptReturnsTrue(t *testing.T) {
	sut, validPasswordHash := makeSut()

	sut.SetPassword("valid_password")
	sut.SetHashedPassword(validPasswordHash)

	passwordIsValid := sut.Compare()

	assert.Equal(t, true, passwordIsValid)
}

func TestShouldReturnFalseIfBCryptReturnsFalse(t *testing.T) {
	sut, validPasswordHash := makeSut()

	sut.SetPassword("invalid_password")
	sut.SetHashedPassword(validPasswordHash)

	passwordIsValid := sut.Compare()

	assert.Equal(t, false, passwordIsValid)
}
