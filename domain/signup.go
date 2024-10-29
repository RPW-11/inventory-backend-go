package domain

type SignupRequest struct {
	FullName    string `json:"fullName" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken string `json:"accessToken"`
}

type SignupUsecase interface {
	Create(user *User) error
	GetUserByEmail(email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (string, error)
	CreateRefreshToken(user *User, secret string, expiry int) (string, error)
}
