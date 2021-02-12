package tokengenerator

// ITokenGenerator interface
type ITokenGenerator interface {
	GetUserID() string
	SetUserID(userID string)
}
