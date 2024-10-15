package usecase

import (
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

func (lu *loginUsecase) GetUserByEmail(email string) (domain.User, error) {
	return lu.userRepository.GetByEmail(email)
}

func (lu *loginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
