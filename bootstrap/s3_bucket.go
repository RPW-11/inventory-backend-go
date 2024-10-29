package bootstrap

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3Session(env *Env) *s3.S3 {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(env.AWS_Region),
	})

	if err != nil {
		log.Fatal("Failed to connect to the S3 bucket: ", err)
	}

	log.Println("Successfully connected to AWS")

	svc := s3.New(sess)

	return svc
}
