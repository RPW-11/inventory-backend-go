package usecase

import (
	"mime/multipart"

	"github.com/RPW-11/inventory_management_be/domain"
)

type userUsecase struct {
	userRepository    domain.UserRepository
	storageRepository domain.StorageRepository
}

func NewUserUsecase(ur domain.UserRepository, sr domain.StorageRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository:    ur,
		storageRepository: sr,
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

func (uu *userUsecase) UpdateProfilePicture(file multipart.File, fileHeader *multipart.FileHeader) error {
	_, err := uu.storageRepository.UploadImage(domain.PROFILE_DIR, file, fileHeader)

	return err
}

func (uu *userUsecase) GetAllUsers() ([]domain.User, error) {
	return uu.userRepository.Fetch()
}
