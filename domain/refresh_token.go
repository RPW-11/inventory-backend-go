package domain

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(id string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractPositionIDFromToken(requestToken string, secret string) (string, string, error)
}
