package domain

import "time"

const (
	TableUser = "User"
)

type User struct {
	ID        string    `gorm:"column:id;primaryKey"`
	FullName  string    `gorm:"column:full_name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	Position  string    `gorm:"column:position"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return TableUser
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}
