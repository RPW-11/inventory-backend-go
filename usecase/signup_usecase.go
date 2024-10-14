package usecase

import "github.com/RPW-11/inventory_management_be/domain"

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
