package luber

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"

// ILoadUserByEmailRepository interface
type ILoadUserByEmailRepository interface {
	GetEmail() string
	SetEmail(email string)

	Load(email string) entities.IUser
}
