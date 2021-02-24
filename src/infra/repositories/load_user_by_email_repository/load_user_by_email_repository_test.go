package loaduserbyemailrepository_test

import (
	"context"
	"testing"

	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func makeSut() (luber.ILoadUserByEmailRepository, *mongo.Collection) {
	// userModel := nil

	loadUserByEmailRepository := luber.NewLoadUserByEmailRepository(nil)

	loadUserByEmailRepository.SetEmail("valid_email@mail.com")

	return loadUserByEmailRepository, nil
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

	defer ctx.Done()
	defer userModel.Drop(ctx)

	result, _ := userModel.InsertOne(ctx, bson.D{
		{
			Key:   "email",
			Value: "valid_email@mail.com",
		},
	})

	user, err := sut.Load()
	insertedUserId := result.InsertedID.(primitive.ObjectID).Hex()

	assert.Equal(t, user.GetPassword(), user.GetPassword())
	assert.Equal(t, insertedUserId, user.GetID())
	assert.Nil(t, err)
}

func TestShouldReturnAnErrorIfEmailIsNotProvided(t *testing.T) {
	sut := luber.NewLoadUserByEmailRepository(nil)

	user, err := sut.Load()

	assert.Equal(t, "Missing param: Email", err.GetError().Error())
	assert.Nil(t, user)
}
