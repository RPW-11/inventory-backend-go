package repository

import (
	"github.com/RPW-11/inventory_management_be/domain"
	"gorm.io/gorm"
)

type productRepository struct {
	database *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		database: db,
	}
}

func (pr *productRepository) Create(product *domain.Product) error {
	result := pr.database.Create(product)
	return result.Error
}

func (pr *productRepository) GetByID(id string) (domain.Product, error) {
	var product domain.Product
	result := pr.database.Where("id = ?", id).First(&product)

	return product, result.Error
}

func (pr *productRepository) Fetch() ([]domain.Product, error) {
	var products []domain.Product
	result := pr.database.Find(&products)

	return products, result.Error
}
