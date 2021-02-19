package loaduserbyemailrepository

import (
	"context"

	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/database"
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()
var collection = database.GetCollection("users")

// LoadUserByEmailRepository struct
type LoadUserByEmailRepository struct {
	Email     string
	User      entities.IUser
	UserModel *mongo.Collection
}

// NewLoadUserByEmailRepository func
func NewLoadUserByEmailRepository(userModel *mongo.Collection) ILoadUserByEmailRepository {
	return &LoadUserByEmailRepository{
		UserModel: collection,
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
	user := entities.NewUser()

	user.SetEmail("victor@mail.com")
	user.SetPassword("lalala")

	_, err := luber.UserModel.InsertOne(ctx, user)

	if err != nil {
		return nil, shared_custom_errors.NewDefaultError("LoadUserByEmailRepository.Load()")
	}

	luber.User = user

	return luber.User, nil
}
