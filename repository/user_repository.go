package repository

import (
	"net/http"

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

func (ur *userRepository) Create(user *domain.User) *domain.CustomError {
	result := ur.database.Create(user)

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (ur *userRepository) GetByEmail(email string) (domain.User, *domain.CustomError) {
	var user domain.User
	result := ur.database.Where("email = ?", email).Find(&user)

	if result.RowsAffected == 0 {
		return user, domain.NewCustomError("user doesn't exists", http.StatusBadRequest)
	}

	if result.Error != nil {
		return user, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return user, nil
}

func (ur *userRepository) GetByID(id string) (domain.User, *domain.CustomError) {
	var user domain.User
	result := ur.database.Where("id = ?", id).Find(&user)

	if result.RowsAffected == 0 {
		return user, domain.NewCustomError("user doesn't exists", http.StatusBadRequest)
	}

	if result.Error != nil {
		return user, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return user, nil
}

func (ur *userRepository) ModifyUserByID(id string, user *domain.User) *domain.CustomError {
	result := ur.database.Model(&domain.User{ID: id}).Updates(user)

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (ur *userRepository) Fetch() ([]domain.User, *domain.CustomError) {
	var users []domain.User
	result := ur.database.Select("id", "email", "position", "full_name", "phone_number", "image_url", "created_at", "updated_at").Find(&users)

	if result.Error != nil {
		return users, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return users, nil
}
