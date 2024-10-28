package domain

import "time"

const (
	TableUser = "User"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	FullName  string    `gorm:"column:full_name" json:"fullName"`
	Email     string    `gorm:"column:email" json:"email"`
	Password  string    `gorm:"column:password" json:"password"`
	Position  string    `gorm:"column:position" json:"position"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

type Profile struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	FullName  string    `gorm:"column:full_name" json:"fullName"`
	Email     string    `gorm:"column:email" json:"email"`
	Position  string    `gorm:"column:position" json:"position"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (User) TableName() string {
	return TableUser
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
	Fetch() ([]User, error)
}

type UserUsecase interface {
	GetProfile(id string) (Profile, error)
	GetAllUsers() ([]User, error)
}
