package usecase

import (
	"mime/multipart"
	"net/http"
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

func (pu *productUsecase) Create(product *domain.Product) *domain.CustomError {
	return pu.productRepository.Create(product)
}

func (pu *productUsecase) GetByID(id string) (domain.Product, *domain.CustomError) {
	return pu.productRepository.GetByID(id)
}

func (pu *productUsecase) Fetch(name string, pageSize, offset int) ([]domain.Product, *domain.CustomError) {
	return pu.productRepository.Fetch(name, pageSize, offset)
}

func (pu *productUsecase) AddProductImages(fileHeaders []*multipart.FileHeader, productId string) *domain.CustomError {
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			return domain.NewCustomError(err.Error(), http.StatusBadRequest)
		}

		defer file.Close()

		contentType := fileHeader.Header.Get("Content-Type")
		fileHeader.Filename = uuid.NewString() + "." + strings.Split(contentType, "/")[1]

		imgUrl, custErr := pu.storageRepository.UploadImage(domain.IMAGE_DIR, file, fileHeader)
		if custErr != nil {
			return custErr
		}

		custErr = pu.productRepository.AddImageUrl(productId, imgUrl)
		if custErr != nil {
			return custErr
		}
	}
	return nil
}

func (pu *productUsecase) DeleteProductImage(productImageId string) *domain.CustomError {
	// check if the product image exists
	productImage, err := pu.productRepository.GetImageById(productImageId)
	if err != nil {
		return err
	}

	// delete from bucket
	fileName := productImage.ImageUrl[strings.LastIndex(productImage.ImageUrl, "/")+1:]
	err = pu.storageRepository.DeleteImage(domain.IMAGE_DIR, fileName)
	if err != nil {
		return err
	}

	// delete from the db
	err = pu.productRepository.DeleteImageUrl(productImageId)

	return err
}

func (pu *productUsecase) DeleteProduct(productId string) *domain.CustomError {
	images, err := pu.productRepository.GetImagesByProductId(productId)
	if err != nil {
		return err
	}

	for _, img := range images {
		fileName := img.ImageUrl[strings.LastIndex(img.ImageUrl, "/")+1:]
		err = pu.storageRepository.DeleteImage(domain.IMAGE_DIR, fileName)
		if err != nil {
			return err
		}
	}

	err = pu.productRepository.DeleteById(productId)
	if err != nil {
		return err
	}

	return nil
}
