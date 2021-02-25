package loaduserbyemailrepository

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	"github.com/go-pg/pg/v10"
)

// LoadUserByEmailRepository struct
type LoadUserByEmailRepository struct {
	Email    string
	User     entities.IUser
	Database *pg.DB
}

// NewLoadUserByEmailRepository func
func NewLoadUserByEmailRepository(db *pg.DB) ILoadUserByEmailRepository {
	return &LoadUserByEmailRepository{
		Database: db,
	}
}

// GetEmail LoadUserByEmailRepository method
func (luber *LoadUserByEmailRepository) GetEmail() string {
	return luber.Email
}

// SetEmail LoadUserByEmailRepository method
func (luber *LoadUserByEmailRepository) SetEmail(email string) {
	luber.Email = email
}

// GetUser LoadUserByEmailRepository method
func (luber *LoadUserByEmailRepository) GetUser() entities.IUser {
	return luber.User
}

// SetUser LoadUserByEmailRepository method
func (luber *LoadUserByEmailRepository) SetUser(user entities.IUser) {
	luber.User = user
}

// Load LoadUserByEmailRepository method
func (luber *LoadUserByEmailRepository) Load() (entities.IUser, shared_custom_errors.IDefaultError) {
	if luber.Email == "" {
		return nil, shared_custom_errors.NewMissingParamError("Email")
	}

	user := entities.NewUser()

	err := luber.Database.Model(user).Where("email = ?", luber.Email).First()

	if err != nil {
		return nil, shared_custom_errors.NewDefaultError("LoadUserByEmailRepository.Load()")
	}

	luber.User = user

	return luber.User, nil
}
