package domain

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}

type LoginUsecase interface {
	GetUserByEmail(email string) (User, *CustomError)
	CreateAccessToken(user *User, secret string, expiry int) (string, *CustomError)
	CreateRefreshToken(user *User, secret string, expiry int) (string, *CustomError)
}
