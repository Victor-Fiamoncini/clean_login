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

	isEmailValid := sut.Run("valid_email@mail.com")

	assert.Equal(t, true, isEmailValid)
}

func TestShouldReturnFalseIfValidatorReturnsFalse(t *testing.T) {
	sut := makeSut()

	isEmailValid := sut.Run("invalid_email")

	assert.Equal(t, false, isEmailValid)
}
