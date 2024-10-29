package domain

import "mime/multipart"

type StorageRepository interface {
	UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (string, error)
	DeleteImage(fileName string) error
}
