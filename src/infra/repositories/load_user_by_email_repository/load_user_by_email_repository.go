package loaduserbyemailrepository

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

// LoadUserByEmailRepository struct
type LoadUserByEmailRepository struct {
	Email     string
	User      entities.IUser
	UserModel *mongo.Collection
}

// NewLoadUserByEmailRepository func
func NewLoadUserByEmailRepository(userModel *mongo.Collection) ILoadUserByEmailRepository {
	return &LoadUserByEmailRepository{
		UserModel: userModel,
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
	return luber.User, nil
}
