package encrypter

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func makeSut() IEncrypter {
	return NewEncrypter()
}

func TestShouldReturnTrueIfBCryptReturnsTrue(t *testing.T) {
	sut := makeSut()

	sut.SetPassword("any_password")
	sut.SetHashedPassword("hashed_password")

	passwordIsValid := sut.Compare()

	assert.Equal(t, true, passwordIsValid)
}
