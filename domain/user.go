package domain

import (
	"mime/multipart"
	"time"
)

const (
	TableUser = "User"
)

type User struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	FullName    string    `gorm:"column:full_name" json:"fullName"`
	Email       string    `gorm:"column:email" json:"email"`
	Password    string    `gorm:"column:password" json:"password"`
	Position    string    `gorm:"column:position" json:"position"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phoneNumber"`
	ImageUrl    string    `gorm:"column:image_url" json:"imgUrl"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

type Profile struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	FullName    string    `gorm:"column:full_name" json:"fullName"`
	Email       string    `gorm:"column:email" json:"email"`
	Position    string    `gorm:"column:position" json:"position"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phoneNumber"`
	ImageUrl    string    `gorm:"column:image_url" json:"imgUrl"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (User) TableName() string {
	return TableUser
}

type UserRepository interface {
	Create(user *User) *CustomError
	GetByEmail(email string) (User, *CustomError)
	GetByID(id string) (User, *CustomError)
	ModifyUserByID(id string, user *User) *CustomError
	Fetch(name string, pageSize, offset int) ([]User, *CustomError)
}

type UserUsecase interface {
	GetProfile(id string) (Profile, *CustomError)
	GetAllUsers(name string, pageSize, offset int) ([]User, *CustomError)
	UpdateProfilePicture(userId string, file multipart.File, fileHeader *multipart.FileHeader) *CustomError
}
