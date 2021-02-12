package mocks

import (
	"github.com/Victor-Fiamoncini/auth_clean_architecture/src/domain/entities"
	luber "github.com/Victor-Fiamoncini/auth_clean_architecture/src/infra/repositories/load_user_by_email_repository"
)

// LoadUserByEmailRepositorySpy struct
type LoadUserByEmailRepositorySpy struct {
	Email string
	User  entities.IUser
}

// NewLoadUserByEmailRepositorySpy func
func NewLoadUserByEmailRepositorySpy() luber.ILoadUserByEmailRepository {
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

// GetUser LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) GetUser() entities.IUser {
	return lubers.User
}

// SetUser LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) SetUser(user entities.IUser) {
	lubers.User = user
}

// Load LoadUserByEmailRepositorySpy method
func (lubers *LoadUserByEmailRepositorySpy) Load(email string) entities.IUser {
	lubers.Email = email

	user := entities.NewUser()

	return user
}
