package domain

type SignupRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SignupResponse struct {
	Payload string
}

type SignupUsecase interface {
	Create(user *User) error
	GetUserByEmail(email string) (User, error)
}
