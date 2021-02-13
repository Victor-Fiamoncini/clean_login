package encrypter

import shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"

// IEncrypter interface
type IEncrypter interface {
	GetPassword() string
	SetPassword(password string)
	GetHashedPassword() string
	SetHashedPassword(hashedPassword string)
	GetIsValid() bool
	SetIsValid(hashedPassword bool)

	Compare() (bool, shared_custom_errors.IDefaultError)
}
