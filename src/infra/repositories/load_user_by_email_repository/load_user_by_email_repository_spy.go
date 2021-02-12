package luber

import "github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"

// LoadUserByEmailRepositorySpy struct
type LoadUserByEmailRepositorySpy struct {
	Email string
	User  entities.IUser
}

// NewLoadUserByEmailRepositorySpy func
func NewLoadUserByEmailRepositorySpy() ILoadUserByEmailRepository {
	return &LoadUserByEmailRepositorySpy{}
}

// GetEmail LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) GetEmail() string {
	return lubers.Email
}

// SetEmail LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) SetEmail(email string) {
	lubers.Email = email
}

// Load LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) Load(email string) entities.IUser {
	lubers.Email = email

	return nil
}
