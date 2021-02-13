package encrypter

import (
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
func (e *Encrypter) Compare() bool {
	err := bcrypt.CompareHashAndPassword([]byte(e.HashedPassword), []byte(e.Password))

	e.IsValid = err == nil

	return e.IsValid
}
