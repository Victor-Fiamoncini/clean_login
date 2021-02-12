package entities

// User struct
type User struct {
	ID       string
	Password string
}

// NewUser func
func NewUser() IUser {
	return &User{}
}

// GetID User method
func (u *User) GetID() string {
	return u.ID
}

// SetID User method
func (u *User) SetID(id string) {
	u.ID = id
}

// GetPassword User method
func (u *User) GetPassword() string {
	return u.Password
}

// SetPassword User method
func (u *User) SetPassword(password string) {
	u.Password = password
}
