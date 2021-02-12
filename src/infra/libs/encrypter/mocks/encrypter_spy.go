package mocks

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/libs/encrypter"

// EncrypterSpy struct
type EncrypterSpy struct {
	Password       string
	HashedPassword string
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
	return es.Password
}

// SetHashedPassword EncrypterSpy method
func (es *EncrypterSpy) SetHashedPassword(password string) {
	es.Password = password
}

// Compare EncrypterSpy method
func (es *EncrypterSpy) Compare(password string, hashedPassword string) {
}
