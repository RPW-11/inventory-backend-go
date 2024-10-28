package repository

import (
	"fmt"

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

func (ur *userRepository) GetByID(id string) (domain.User, error) {
	var user domain.User
	result := ur.database.Where("id = ?", id).Find(&user)

	if result.RowsAffected == 0 {
		return user, fmt.Errorf("no user found")
	}

	return user, result.Error
}

func (ur *userRepository) Fetch() ([]domain.User, error) {
	var users []domain.User
	result := ur.database.Select("id", "email", "position", "full_name", "created_at", "updated_at").Find(&users)

	return users, result.Error
}
