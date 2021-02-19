package loaduserbyemailrepository_test

import (
	"context"
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func makeSut() (luber.ILoadUserByEmailRepository, *mongo.Collection) {
	userModel := database.GetCollection("users")

	loadUserByEmailRepository := luber.NewLoadUserByEmailRepository(userModel)

	loadUserByEmailRepository.SetEmail("valid_email@mail.com")

	return loadUserByEmailRepository, userModel
}

func TestShouldReturnNullAndAnErrorIfNoUserIsFound(t *testing.T) {
	sut, _ := makeSut()

	sut.SetEmail("invalid_email@mail.com")

	user, err := sut.Load()

	assert.Nil(t, user)
	assert.Equal(t, "Error with: LoadUserByEmailRepository.Load()", err.GetError().Error())
}

func TestShouldReturnAnUserIfUserIsFound(t *testing.T) {
	sut, userModel := makeSut()
	ctx := context.Background()

	defer userModel.Drop(ctx)
	defer ctx.Done()

	userModel.InsertOne(ctx, bson.D{
		{
			Key:   "email",
			Value: "valid_email@mail.com",
		},
	})

	user, err := sut.Load()

	assert.Equal(t, "valid_email@mail.com", user.GetEmail())
	assert.Nil(t, err)
}
