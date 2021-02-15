package loaduserbyemailrepository_test

import (
	"context"
	"log"
	"testing"
	"time"

	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConnection() (*mongo.Client, context.Context, *mongo.Collection) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, connectErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo"))

	if connectErr != nil {
		log.Fatal(connectErr)
	}

	userModel := mongoClient.Database("auth_clean_architecture").Collection("users")

	return mongoClient, ctx, userModel
}

func makeSut() (luber.ILoadUserByEmailRepository, *mongo.Client, context.Context, *mongo.Collection) {
	mongoClient, ctx, userModel := mongoConnection()

	loadUserByEmailRepository := luber.NewLoadUserByEmailRepository(userModel)

	loadUserByEmailRepository.SetEmail("valid_email@mail.com")

	return loadUserByEmailRepository, mongoClient, ctx, userModel
}

func TestShouldReturnNullAndAnErrorIfNoUserIsFound(t *testing.T) {
	sut, mongoClient, ctx, userModel := makeSut()

	defer userModel.Drop(ctx)
	defer mongoClient.Disconnect(ctx)
	defer ctx.Done()

	sut.SetEmail("invalid_email@mail.com")

	user, err := sut.Load()

	assert.Nil(t, user)
	assert.Equal(t, "Error with: LoadUserByEmailRepository.Load()", err.GetError().Error())
}

func TestShouldReturnAnUserIfUserIsFound(t *testing.T) {
	sut, mongoClient, ctx, userModel := makeSut()

	defer userModel.Drop(ctx)
	defer mongoClient.Disconnect(ctx)
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
