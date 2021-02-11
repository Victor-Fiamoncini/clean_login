package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnTrueIfValidatorReturnsTrue(t *testing.T) {
	sut := NewEmailValidator()

	isEmailValid := sut.Run("valid_email@mail.com")

	assert.Equal(t, true, isEmailValid)
}

func TestShouldReturnFalseIfValidatorReturnsFalse(t *testing.T) {
	sut := NewEmailValidator()

	isEmailValid := sut.Run("invalid_email")

	assert.Equal(t, false, isEmailValid)
}
