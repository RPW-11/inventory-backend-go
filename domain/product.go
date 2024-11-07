package domain

import (
	"mime/multipart"
	"time"
)

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
	ProductID string    `gorm:"column:product_id" json:"productId"`
	ImageUrl  string    `gorm:"column:image_url" json:"imageUrl"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
}

func (ProductImage) TableName() string {
	return TableProductImage
}

type ProductRepository interface {
	Create(product *Product) *CustomError
	GetByID(id string) (Product, *CustomError)
	Fetch(name string, pageSize, offset int) ([]Product, *CustomError)
	AddImageUrl(productId, imgUrl string) *CustomError
	DeleteImageUrl(productImageId string) *CustomError
	DeleteById(productId string) *CustomError
	GetImageById(productImageId string) (ProductImage, *CustomError)
	GetImagesByProductId(productId string) ([]ProductImage, *CustomError)
}

type ProductUsecase interface {
	Create(product *Product) *CustomError
	GetByID(id string) (Product, *CustomError)
	Fetch(name string, pageSize, offset int) ([]Product, *CustomError)
	AddProductImages(fileHeaders []*multipart.FileHeader, productId string) *CustomError
	DeleteProductImage(productImageId string) *CustomError
	DeleteProduct(productId string) *CustomError
}
