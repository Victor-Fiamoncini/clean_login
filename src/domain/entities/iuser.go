package entities

// IUser interface
type IUser interface {
	GetID() string
	SetID(id string)
	GetEmail() string
	SetEmail(email string)
	GetPassword() string
	SetPassword(password string)
	GetAccessToken() string
	SetAccessToken(accessToken string)
}
