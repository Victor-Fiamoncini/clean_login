package validators_test

import (
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/validators"
	"github.com/stretchr/testify/assert"
)

func makeSut() validators.IEmailValidator {
	emailValidator := validators.NewEmailValidator()

	return emailValidator
}

func TestShouldReturnTrueIfValidatorReturnsTrue(t *testing.T) {
	sut := makeSut()

	sut.SetEmail("valid_email@mail.com")

	isEmailValid := sut.Run()

	assert.Equal(t, true, isEmailValid)
}

func TestShouldReturnFalseIfValidatorReturnsFalse(t *testing.T) {
	sut := makeSut()

	sut.SetEmail("invalid_email")

	isEmailValid := sut.Run()

	assert.Equal(t, false, isEmailValid)
}

func TestShouldReturnFalseIfNoEmailIsProvided(t *testing.T) {
	sut := makeSut()

	isEmailValid := sut.Run()

	assert.Equal(t, false, isEmailValid)
}
