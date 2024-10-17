package usecase

import (
	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/RPW-11/inventory_management_be/internal/tokenutil"
)

type refreshTokenUsecase struct {
	UserRepository domain.UserRepository
}

func NewRefreshTokenUsecase(ur domain.UserRepository) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		UserRepository: ur,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(id string) (domain.User, error) {
	return rtu.UserRepository.GetByID(id)

}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractPositionIDFromToken(requestToken string, secret string) (string, string, error) {
	return tokenutil.ExtractPositionIDFromToken(requestToken, secret)
}
