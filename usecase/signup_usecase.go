package usecase

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/RPW-11/inventory_management_be/internal/tokenutil"
)

type signupUsecase struct {
	userRepository domain.UserRepository
}

func NewSignupUsecase(ur domain.UserRepository) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: ur,
	}
}

func (su *signupUsecase) Create(user *domain.User) *domain.CustomError {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) GetUserByEmail(email string) (domain.User, *domain.CustomError) {
	return su.userRepository.GetByEmail(email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, *domain.CustomError) {
	token, err := tokenutil.CreateAccessToken(user, secret, expiry)
	if err != nil {
		return token, domain.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return token, nil
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, *domain.CustomError) {
	token, err := tokenutil.CreateRefreshToken(user, secret, expiry)
	if err != nil {
		return token, domain.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return token, nil
}
