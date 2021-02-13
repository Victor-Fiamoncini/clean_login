package encrypter

import (
	shared_custom_errors "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/errors"
	"golang.org/x/crypto/bcrypt"
)

// Encrypter struct
type Encrypter struct {
	Password       string
	HashedPassword string
	IsValid        bool
}

// NewEncrypter func
func NewEncrypter() IEncrypter {
	return &Encrypter{}
}

// GetPassword Encrypter method
func (e *Encrypter) GetPassword() string {
	return e.Password
}

// SetPassword Encrypter method
func (e *Encrypter) SetPassword(password string) {
	e.Password = password
}

// GetHashedPassword Encrypter method
func (e *Encrypter) GetHashedPassword() string {
	return e.HashedPassword
}

// SetHashedPassword Encrypter method
func (e *Encrypter) SetHashedPassword(hashedPassword string) {
	e.HashedPassword = hashedPassword
}

// GetIsValid Encrypter method
func (e *Encrypter) GetIsValid() bool {
	return e.IsValid
}

// SetIsValid Encrypter method
func (e *Encrypter) SetIsValid(isValid bool) {
	e.IsValid = isValid
}

// Compare Encrypter method
func (e *Encrypter) Compare() (bool, shared_custom_errors.IDefaultError) {
	if e.Password == "" {
		return false, shared_custom_errors.NewMissingParamError("Password")
	}

	if e.HashedPassword == "" {
		return false, shared_custom_errors.NewMissingParamError("HashedPassword")
	}

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(e.HashedPassword), []byte(e.Password))

	if bcryptErr != nil {
		e.IsValid = false

		return false, shared_custom_errors.NewUnexpectedError("Encrypter.Compare()")
	}

	e.IsValid = true

	return e.IsValid, nil
}
