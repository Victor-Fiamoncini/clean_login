package loaduserbyemailrepository_test

import (
	"testing"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/assert"
)

func makeSut() (luber.ILoadUserByEmailRepository, *pg.DB) {
	db := database.OpenConnection()

	loadUserByEmailRepository := luber.NewLoadUserByEmailRepository(db)

	loadUserByEmailRepository.SetEmail("valid_email@mail.com")

	return loadUserByEmailRepository, db
}

func TestShouldReturnNullAndAnErrorIfNoUserIsFound(t *testing.T) {
	sut, _ := makeSut()

	sut.SetEmail("invalid_email@mail.com")

	user, err := sut.Load()

	assert.Nil(t, user)
	assert.Equal(t, "Error with: LoadUserByEmailRepository.Load()", err.GetError().Error())
}

func TestShouldReturnAnUserIfUserIsFound(t *testing.T) {
	sut, db := makeSut()

	newUser := entities.NewUser()

	newUser.SetEmail("valid_email@mail.com")
	newUser.SetPassword("valid_password")

	db.Model(newUser).Insert()

	user, err := sut.Load()

	assert.Equal(t, user.GetEmail(), "valid_email@mail.com")
	assert.Equal(t, user.GetPassword(), "valid_password")
	assert.Nil(t, err)

	db.Model(newUser).Where("email = ?", newUser.GetEmail()).Delete()
}

func TestShouldReturnAnErrorIfEmailIsNotProvided(t *testing.T) {
	sut := luber.NewLoadUserByEmailRepository(nil)

	user, err := sut.Load()

	assert.Equal(t, "Missing param: Email", err.GetError().Error())
	assert.Nil(t, user)
}
