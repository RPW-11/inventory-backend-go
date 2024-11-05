package usecase

import (
	"net/http"

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

func (rtu *refreshTokenUsecase) GetUserByID(id string) (domain.User, *domain.CustomError) {
	return rtu.UserRepository.GetByID(id)

}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, *domain.CustomError) {
	token, err := tokenutil.CreateAccessToken(user, secret, expiry)
	if err != nil {
		return token, domain.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return token, nil
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, *domain.CustomError) {
	token, err := tokenutil.CreateRefreshToken(user, secret, expiry)
	if err != nil {
		return token, domain.NewCustomError(err.Error(), http.StatusInternalServerError)
	}

	return token, nil
}

func (rtu *refreshTokenUsecase) ExtractPositionIDFromToken(requestToken string, secret string) (string, string, *domain.CustomError) {
	id, position, err := tokenutil.ExtractPositionIDFromToken(requestToken, secret)
	if err != nil {
		return id, position, domain.NewCustomError(err.Error(), http.StatusBadRequest)
	}

	return id, position, nil
}
