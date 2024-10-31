package domain

import "time"

const (
	TableProduct      = "Product"
	TableProductImage = "ProductImage"
)

type Product struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	Price       float64   `gorm:"column:price" json:"price"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Product) TableName() string {
	return TableProduct
}

type ProductImage struct {
	ID        int       `gorm:"column:id;primaryKey" json:"id"`
	ProductID string    `gorm:"column:id" json:"productId"`
	ImageUrl  string    `gorm:"column:image_url" json:"imageUrl"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (ProductImage) TableName() string {
	return TableProductImage
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id string) (Product, error)
	Fetch(name string) ([]Product, error)
	AddImageUrl(productId, imgUrl string) error
}

type ProductUsecase interface {
	Create(product *Product) error
	GetByID(id string) (Product, error)
	Fetch(name string) ([]Product, error)
}
