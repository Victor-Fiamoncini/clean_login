package updateaccesstokenrepository

// IUpdateAccessTokenRepository interface
type IUpdateAccessTokenRepository interface {
	GetUserID() string
	SetUserID(userID string)
	GetAccessToken() string
	SetAccessToken(accessToken string)

	Update()
}
