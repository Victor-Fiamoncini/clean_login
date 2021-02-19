package updateaccesstokenrepository_test

import (
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"
	uatr "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

func makeSut() (uatr.IUpdateAccessTokenRepository, *mongo.Collection) {
	userModel := database.GetCollection("users")

	updateAccessTokenRepository := uatr.NewUpdateAccessTokenRepository()

	updateAccessTokenRepository.SetUserID("valid_id")
	updateAccessTokenRepository.SetAccessToken("valid_token")

	return updateAccessTokenRepository, userModel
}

func TestShouldUpdateTheUserWithTheGivenAccessToken(t *testing.T) {
	sut, _ := makeSut()

	sut.Update()

	assert.Equal(t, user, sut.GetAccessToken())
}
