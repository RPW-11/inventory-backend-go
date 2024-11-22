package repository

import (
	"net/http"
	"time"

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

func (pr *productRepository) Create(product *domain.Product) *domain.CustomError {
	result := pr.database.Create(product)

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (pr *productRepository) GetByID(id string) (domain.Product, *domain.CustomError) {
	var product domain.Product
	result := pr.database.Where("id = ?", id).Find(&product)

	if result.RowsAffected == 0 {
		return product, domain.NewCustomError("product doesn't exists", http.StatusBadRequest)
	}

	if result.Error != nil {
		return product, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return product, nil
}

func (pr *productRepository) Fetch(name string, pageSize, offset int) ([]domain.Product, *domain.CustomError) {
	var products []domain.Product
	query := pr.database.Model(&products)

	query.Order("updated_at desc")

	if name != "" {
		query.Where("lower(name) LIKE lower(?)", "%"+name+"%")
	}

	result := query.Limit(pageSize).Offset(offset).Find(&products)

	if result.Error != nil {
		return products, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return products, nil
}

func (pr *productRepository) AddImageUrl(productId, imgUrl string) *domain.CustomError {
	productImage := domain.ProductImage{
		ProductID: productId,
		ImageUrl:  imgUrl,
	}

	result := pr.database.Create(&productImage)

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (pr *productRepository) GetImageById(productImageId string) (domain.ProductImage, *domain.CustomError) {
	var productImage domain.ProductImage
	result := pr.database.Where("id = ?", productImageId).Find(&productImage)

	if result.RowsAffected == 0 {
		return productImage, domain.NewCustomError("product's image(s) doesn't exist", http.StatusBadRequest)
	}

	if result.Error != nil {
		return productImage, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return productImage, nil
}

func (pr *productRepository) GetImagesByProductId(productId string) ([]domain.ProductImage, *domain.CustomError) {
	var productImages []domain.ProductImage
	result := pr.database.Where("product_id = ?", productId).Find(&productImages)

	if result.Error != nil {
		return productImages, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return productImages, nil
}

func (pr *productRepository) DeleteImageUrl(productImageId string) *domain.CustomError {
	result := pr.database.Delete(&domain.ProductImage{}, "id = ?", productImageId)

	if result.RowsAffected == 0 {
		return domain.NewCustomError("invalid product image id", http.StatusBadRequest)
	}

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (pr *productRepository) DeleteById(productId string) *domain.CustomError {
	result := pr.database.Delete(&domain.Product{}, "id = ?", productId)

	if result.RowsAffected == 0 {
		return domain.NewCustomError("invalid product id", http.StatusBadRequest)
	}

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (pr *productRepository) ModifyByID(product *domain.Product) *domain.CustomError {
	product.UpdatedAt = time.Now()
	result := pr.database.Model(&domain.Product{ID: product.ID}).Updates(product)

	if result.RowsAffected == 0 {
		return domain.NewCustomError("invalid product id", http.StatusBadRequest)
	}

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}
