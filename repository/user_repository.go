package repository

import (
	"github.com/RPW-11/inventory_management_be/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	database *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		database: db,
	}
}

func (ur *userRepository) Create(user *domain.User) error {
	result := ur.database.Create(user)

	return result.Error
}

func (ur *userRepository) GetByEmail(email string) (domain.User, error) {
	var user domain.User
	result := ur.database.Where("email = ?", email).First(&user)

	return user, result.Error
}
