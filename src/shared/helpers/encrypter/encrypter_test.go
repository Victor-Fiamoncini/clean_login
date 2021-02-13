package encrypter_test

import (
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/encrypter"
	"github.com/go-playground/assert/v2"
)

func makeSut() (encrypter.IEncrypter, string) {
	const validPasswordHash = "$2a$09$ekW4Ggqbaixn3x/qyzcreex5RoFIaPAMPLrDaxDZg0QE0Qe008nOW"

	return encrypter.NewEncrypter(), validPasswordHash
}

func TestShouldReturnTrueIfBCryptReturnsTrue(t *testing.T) {
	sut, validPasswordHash := makeSut()

	sut.SetPassword("valid_password")
	sut.SetHashedPassword(validPasswordHash)

	passwordIsValid, _ := sut.Compare()

	assert.Equal(t, true, passwordIsValid)
}

func TestShouldReturnFalseAndAnErrorIfBCryptReturnsFalse(t *testing.T) {
	sut, validPasswordHash := makeSut()

	sut.SetPassword("invalid_password")
	sut.SetHashedPassword(validPasswordHash)

	passwordIsValid, err := sut.Compare()

	assert.Equal(t, false, passwordIsValid)
	assert.Equal(t, "Unexpected Error with: Encrypter.Compare()", err.GetError().Error())
}

func TestShouldReturnFalseAndAnErrorIfPasswordIsNotProvided(t *testing.T) {
	sut, validPasswordHash := makeSut()

	sut.SetHashedPassword(validPasswordHash)

	passwordIsValid, err := sut.Compare()

	assert.Equal(t, false, passwordIsValid)
	assert.Equal(t, "Missing param: Password", err.GetError().Error())
}

func TestShouldReturnFalseAndAnErrorIfHashedPasswordIsNotProvided(t *testing.T) {
	sut, _ := makeSut()

	sut.SetPassword("any_password")

	passwordIsValid, err := sut.Compare()

	assert.Equal(t, false, passwordIsValid)
	assert.Equal(t, "Missing param: HashedPassword", err.GetError().Error())
}
