package usecase

import (
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

func (su *signupUsecase) Create(user *domain.User) error {
	return su.userRepository.Create(user)
}

func (su *signupUsecase) GetUserByEmail(email string) (domain.User, error) {
	return su.userRepository.GetByEmail(email)
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
