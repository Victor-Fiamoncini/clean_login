package loaduserbyemailrepository

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"

// ILoadUserByEmailRepository interface
type ILoadUserByEmailRepository interface {
	GetEmail() string
	SetEmail(email string)
	GetUser() entities.IUser
	SetUser(user entities.IUser)

	Load() entities.IUser
}
