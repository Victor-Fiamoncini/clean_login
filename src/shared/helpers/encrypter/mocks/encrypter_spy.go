package mocks

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/shared/helpers/encrypter"

// EncrypterSpy struct
type EncrypterSpy struct {
	Password       string
	HashedPassword string
	IsValid        bool
}

// NewEncrypterSpy func
func NewEncrypterSpy() encrypter.IEncrypter {
	return &EncrypterSpy{}
}

// GetPassword EncrypterSpy method
func (es *EncrypterSpy) GetPassword() string {
	return es.Password
}

// SetPassword EncrypterSpy method
func (es *EncrypterSpy) SetPassword(password string) {
	es.Password = password
}

// GetHashedPassword EncrypterSpy method
func (es *EncrypterSpy) GetHashedPassword() string {
	return es.HashedPassword
}

// SetHashedPassword EncrypterSpy method
func (es *EncrypterSpy) SetHashedPassword(hashedPassword string) {
	es.HashedPassword = hashedPassword
}

// GetIsValid EncrypterSpy method
func (es *EncrypterSpy) GetIsValid() bool {
	return es.IsValid
}

// SetIsValid EncrypterSpy method
func (es *EncrypterSpy) SetIsValid(isValid bool) {
	es.IsValid = isValid
}

// Compare EncrypterSpy method
func (es *EncrypterSpy) Compare() bool {
	return es.IsValid
}
