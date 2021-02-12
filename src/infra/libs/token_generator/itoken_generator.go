package tokengenerator

// ITokenGenerator interface
type ITokenGenerator interface {
	GetUserID() string
	SetUserID(userID string)
	GetAccessToken() string
	SetAccessToken(accessToken string)

	Generate(userID string) string
}
