package usecase

import "github.com/RPW-11/inventory_management_be/domain"

type productUsecase struct {
	productRepository domain.ProductRepository
}

func NewProductUsecase(pr domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{
		productRepository: pr,
	}
}

func (pu *productUsecase) Create(product *domain.Product) error {
	return pu.productRepository.Create(product)
}

func (pu *productUsecase) GetByID(id string) (domain.Product, error) {
	return pu.productRepository.GetByID(id)
}

func (pu *productUsecase) Fetch() ([]domain.Product, error) {
	return pu.productRepository.Fetch()
}
