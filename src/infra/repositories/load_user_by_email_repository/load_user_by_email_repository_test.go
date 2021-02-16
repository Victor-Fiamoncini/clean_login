package loaduserbyemailrepository_test

import (
	"testing"

	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	"github.com/stretchr/testify/assert"
)

func makeSut() luber.ILoadUserByEmailRepository {
	loadUserByEmailRepository := luber.NewLoadUserByEmailRepository(nil)

	loadUserByEmailRepository.SetEmail("valid_email@mail.com")

	return loadUserByEmailRepository
}

func TestShouldReturnNullAndAnErrorIfNoUserIsFound(t *testing.T) {
	sut := makeSut()

	sut.SetEmail("invalid_email@mail.com")

	user, err := sut.Load()

	assert.Nil(t, user)
	assert.Equal(t, "Error with: LoadUserByEmailRepository.Load()", err.GetError().Error())
}

func TestShouldReturnAnUserIfUserIsFound(t *testing.T) {
	sut := makeSut()

	// userModel.InsertOne(ctx, bson.D{
	// 	{
	// 		Key:   "email",
	// 		Value: "valid_email@mail.com",
	// 	},
	// })

	user, err := sut.Load()

	assert.Equal(t, "valid_email@mail.com", user.GetEmail())
	assert.Nil(t, err)
}
