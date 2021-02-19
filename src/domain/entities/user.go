package entities

// User struct
type User struct {
	ID       string `bson:"_id" json:"_id"`
	Email    string `bson:"email" json:"email"`
	Password string `bson:"password" json:"password"`
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

// GetEmail User method
func (u *User) GetEmail() string {
	return u.Email
}

// SetEmail User method
func (u *User) SetEmail(email string) {
	u.Email = email
}

// GetPassword User method
func (u *User) GetPassword() string {
	return u.Password
}

// SetPassword User method
func (u *User) SetPassword(password string) {
	u.Password = password
}
