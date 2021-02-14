package loaduserbyemailrepository_test

import (
	"context"
	"log"
	"testing"

	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()
var userModel *mongo.Collection

func init() {
	mongoClient, connectErr := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo"))

	if connectErr != nil {
		log.Fatal(connectErr)
	}

	pingErr := mongoClient.Ping(ctx, nil)

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	userModel = mongoClient.Database("auth_clean_architecture").Collection("users")
}

func makeSut() luber.ILoadUserByEmailRepository {
	loadUserByEmailRepository := luber.NewLoadUserByEmailRepository(userModel)

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
