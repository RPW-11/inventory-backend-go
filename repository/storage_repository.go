package repository

import (
	"mime/multipart"

	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type storageRepository struct {
	env     *bootstrap.Env
	storage *s3.S3
}

func NewStorageRepository(env *bootstrap.Env, storage *s3.S3) domain.StorageRepository {
	return &storageRepository{
		env:     env,
		storage: storage,
	}
}

func (sr *storageRepository) UploadImage(dir string, file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	key := dir + fileHeader.Filename
	input := &s3.PutObjectInput{
		Bucket:        aws.String(sr.env.S3Bucket),
		Key:           aws.String(key),
		Body:          file,
		ContentType:   aws.String(fileHeader.Header.Get("Content-Type")),
		ContentLength: &fileHeader.Size,
	}

	_, err := sr.storage.PutObject(input)
	if err != nil {
		return "", err
	}

	return key, nil
}

func (sr *storageRepository) DeleteImage(fileName string) error {
	return nil
}
