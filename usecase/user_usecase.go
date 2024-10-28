package usecase

import "github.com/RPW-11/inventory_management_be/domain"

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (uu *userUsecase) GetProfile(id string) (domain.Profile, error) {
	user, err := uu.userRepository.GetByID(id)
	profile := domain.Profile{
		ID:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Position:  user.Position,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return profile, err
}

func (uu *userUsecase) GetAllUsers() ([]domain.User, error) {
	return uu.userRepository.Fetch()
}
