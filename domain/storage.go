package domain

import "mime/multipart"

const (
	IMAGE_DIR   = "product-images/"
	PROFILE_DIR = "profile-images/"
)

type StorageRepository interface {
	UploadImage(dir string, file multipart.File, fileHeader *multipart.FileHeader) (string, *CustomError)
	DeleteImage(dir, fileName string) *CustomError
}
