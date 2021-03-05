package updateaccesstokenrepository_test

import (
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"
	uatr "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/update_access_token_repository"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func makeSut() (uatr.IUpdateAccessTokenRepository, *pg.DB) {
	db := database.OpenConnection()

	updateAccessTokenRepository := uatr.NewUpdateAccessTokenRepository(db)

	return updateAccessTokenRepository, db
}

func TestShouldUpdateTheUserWithTheGivenAccessToken(t *testing.T) {
	sut, db := makeSut()

	newUser := entities.NewUser()
	newUser.SetEmail("valid_email@mail.com")
	newUser.SetPassword("hashed_password")

	db.Model(newUser).Insert()
	db.Model(newUser).Where("email = ?", "valid_email@mail.com").First()

	insertedUserID := newUser.GetID()

	sut.SetUserID(insertedUserID)
	sut.SetAccessToken("valid_token")

	err := sut.Update()

	updatedUser := entities.NewUser()
	updatedUser.SetID(insertedUserID)

	db.Model(updatedUser).WherePK().Select()

	assert.Equal(t, "valid_token", updatedUser.GetAccessToken())
	assert.Equal(t, insertedUserID, updatedUser.GetID())
	assert.Nil(t, err)

	db.Model(newUser).Where("email = ?", newUser.GetEmail()).Delete()
}

func TestShouldReturnAnErrorIfNoAccessTokenIsProvided(t *testing.T) {
	sut, _ := makeSut()

	sut.SetUserID("any_id")

	err := sut.Update()

	assert.Equal(t, "Missing param: AccessToken", err.GetError().Error())
}

func TestShouldReturnAnErrorIfNoUserIdIsProvided(t *testing.T) {
	sut, _ := makeSut()

	sut.SetAccessToken("any_token")

	err := sut.Update()

	assert.Equal(t, "Missing param: UserID", err.GetError().Error())
}

func TestShouldReturnAnErrorIfUserWasNotFound(t *testing.T) {
	sut, db := makeSut()

	newUser := entities.NewUser()
	newUser.SetEmail("valid_email@mail.com")
	newUser.SetPassword("hashed_password")
	db.Model(newUser).Where("email = ?", newUser.GetEmail()).Delete()

	db.Model(newUser).Insert()

	randomID, _ := uuid.NewUUID()

	sut.SetUserID(randomID.String())
	sut.SetAccessToken("valid_token")

	err := sut.Update()

	assert.Equal(t, "Error with: UpdateAccessTokenRepository.Update()", err.GetError().Error())

	db.Model(newUser).Where("email = ?", newUser.GetEmail()).Delete()
}
