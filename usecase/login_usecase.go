package usecase

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/RPW-11/inventory_management_be/internal/tokenutil"
)

type loginUsecase struct {
	userRepository domain.UserRepository
}

func NewLoginUsecase(ur domain.UserRepository) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: ur,
	}
}

func (lu *loginUsecase) GetUserByEmail(email string) (domain.User, *domain.CustomError) {
	return lu.userRepository.GetByEmail(email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, *domain.CustomError) {
	token, err := tokenutil.CreateAccessToken(user, secret, expiry)
	if err != nil {
		return token, domain.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return token, nil
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, *domain.CustomError) {
	token, err := tokenutil.CreateRefreshToken(user, secret, expiry)
	if err != nil {
		return token, domain.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return token, nil
}
