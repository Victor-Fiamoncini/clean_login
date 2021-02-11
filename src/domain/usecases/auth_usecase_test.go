package usecases_test

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/usecases"
)

func makeSut() usecases.IAuthUseCase {
	return usecases.NewAuthUseCase()
}

// func TestShouldReturnNullIfNoEmailIsProvided(t *testing.T) {
// 	sut := makeSut()
// }
