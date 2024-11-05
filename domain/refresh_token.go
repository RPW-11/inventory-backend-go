package domain

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(id string) (User, *CustomError)
	CreateAccessToken(user *User, secret string, expiry int) (string, *CustomError)
	CreateRefreshToken(user *User, secret string, expiry int) (string, *CustomError)
	ExtractPositionIDFromToken(requestToken string, secret string) (string, string, *CustomError)
}
