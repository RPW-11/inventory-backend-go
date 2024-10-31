package repository

import (
	"fmt"

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
	result := pr.database.Where("id = ?", id).Find(&product)

	if result.RowsAffected == 0 {
		return product, fmt.Errorf("no existing product")
	}

	return product, result.Error
}

func (pr *productRepository) Fetch(name string) ([]domain.Product, error) {
	var products []domain.Product
	query := pr.database.Model(&products)

	if name != "" {
		query.Where("lower(name) LIKE lower(?)", "%"+name+"%")
	}

	result := query.Find(&products)

	return products, result.Error
}

func (pr *productRepository) AddImageUrl(productId, imgUrl string) error {
	productImage := domain.ProductImage{
		ProductID: productId,
		ImageUrl:  imgUrl,
	}

	result := pr.database.Create(&productImage)

	return result.Error
}

func (pr *productRepository) GetImageById(productImageId string) (domain.ProductImage, error) {
	var productImage domain.ProductImage
	result := pr.database.Where("id = ?", productImageId).Find(&productImage)

	if result.RowsAffected == 0 {
		return productImage, fmt.Errorf("no existing product's image")
	}

	return productImage, result.Error
}

func (pr *productRepository) GetImagesByProductId(productId string) ([]domain.ProductImage, error) {
	var productImages []domain.ProductImage
	result := pr.database.Where("product_id = ?", productId).Find(&productImages)

	return productImages, result.Error
}

func (pr *productRepository) DeleteImageUrl(productImageId string) error {
	result := pr.database.Delete(&domain.ProductImage{}, "id = ?", productImageId)

	return result.Error
}
