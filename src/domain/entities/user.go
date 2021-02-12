package entities

// User struct
type User struct {
	Password string
}

// NewUser func
func NewUser() IUser {
	return &User{}
}

// GetPassword User method
func (u *User) GetPassword() string {
	return u.Password
}

// SetPassword User method
func (u *User) SetPassword(password string) {
	u.Password = password
}
