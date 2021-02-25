package entities

import "time"

// User struct
type User struct {
	tableName   struct{}  `pg:"users"`
	ID          string    `json:"id" pg:",pk,type:uuid,default:gen_random_uuid()"`
	Email       string    `json:"email" pg:"type:varchar(255),unique,notnull"`
	Password    string    `json:"password" pg:"type:varchar(255),notnull"`
	AccessToken string    `json:"access_token"`
	CreatedAt   time.Time `json:"created_at" pg:"default:now(),notnull"`
	UpdatedAt   time.Time `json:"updated_at" pg:"default:now(),notnull"`
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

// GetAccessToken User method
func (u *User) GetAccessToken() string {
	return u.AccessToken
}

// SetAccessToken User method
func (u *User) SetAccessToken(accessToken string) {
	u.AccessToken = accessToken
}
