package updateaccesstokenrepository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"
	uatr "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func makeSut() (uatr.IUpdateAccessTokenRepository, *mongo.Collection) {
	userModel := database.GetCollection("users")

	updateAccessTokenRepository := uatr.NewUpdateAccessTokenRepository(userModel)

	return updateAccessTokenRepository, userModel
}

func TestShouldUpdateTheUserWithTheGivenAccessToken(t *testing.T) {
	sut, userModel := makeSut()
	ctx := context.Background()

	defer ctx.Done()
	// defer userModel.Drop(ctx)

	user := entities.NewUser()

	result, _ := userModel.InsertOne(ctx, bson.D{
		{
			Key:   "email",
			Value: "valid_email@mail.com",
		},
		{
			Key:   "password",
			Value: "hashed_password",
		},
		{
			Key:   "access_token",
			Value: "",
		},
	})

	insertedUserId := result.InsertedID.(primitive.ObjectID).Hex()

	sut.SetUserID(insertedUserId)
	sut.SetAccessToken("valid_token")

	sut.Update()

	updatedUser := userModel.FindOne(ctx, bson.D{{
		Key:   "_id",
		Value: insertedUserId,
	}})

	fmt.Println("UPDATEEEEEEEEEEEEEEEE", updatedUser)

	err := updatedUser.Decode(user)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, "valid_token", user.GetAccessToken())
}
