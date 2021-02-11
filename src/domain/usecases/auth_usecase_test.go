package usecases_test

import (
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases"
)

func makeSut() usecases.IAuthUseCase {
	return usecases.NewAuthUseCase()
}

func TestShouldCallLoadUserByEmailRepositoryWithCorrectEmail(t *testing.T) {
	sut := makeSut()
}
