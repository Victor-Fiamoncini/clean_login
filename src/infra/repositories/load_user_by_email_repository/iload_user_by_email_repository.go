package loaduserbyemailrepository

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
)

// ILoadUserByEmailRepository interface
type ILoadUserByEmailRepository interface {
	GetEmail() string
	SetEmail(email string)
	GetUser() entities.IUser
	SetUser(user entities.IUser)

	Load() (entities.IUser, shared_custom_errors.IDefaultError)
}
