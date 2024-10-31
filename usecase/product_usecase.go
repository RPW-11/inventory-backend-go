package usecase

import (
	"mime/multipart"
	"strings"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/google/uuid"
)

type productUsecase struct {
	productRepository domain.ProductRepository
	storageRepository domain.StorageRepository
}

func NewProductUsecase(pr domain.ProductRepository, sr domain.StorageRepository) domain.ProductUsecase {
	return &productUsecase{
		productRepository: pr,
		storageRepository: sr,
	}
}

func (pu *productUsecase) Create(product *domain.Product) error {
	return pu.productRepository.Create(product)
}

func (pu *productUsecase) GetByID(id string) (domain.Product, error) {
	return pu.productRepository.GetByID(id)
}

func (pu *productUsecase) Fetch(name string) ([]domain.Product, error) {
	return pu.productRepository.Fetch(name)
}

func (pu *productUsecase) AddProductImages(fileHeaders []*multipart.FileHeader, productId string) error {
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			return err
		}

		defer file.Close()

		contentType := fileHeader.Header.Get("Content-Type")
		fileHeader.Filename = uuid.NewString() + "." + strings.Split(contentType, "/")[1]
		imgUrl, err := pu.storageRepository.UploadImage(domain.IMAGE_DIR, file, fileHeader)
		if err != nil {
			return err
		}

		err = pu.productRepository.AddImageUrl(productId, imgUrl)
		if err != nil {
			return err
		}
	}
	return nil
}
