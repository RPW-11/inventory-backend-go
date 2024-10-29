package bootstrap

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
)

type Application struct {
	Env *Env
	Db  *gorm.DB
	SVC *s3.S3
}

func App() Application {
	app := Application{}
	app.Env = NewEnv()
	app.Db = NewPostgresConnection(app.Env)
	app.SVC = NewS3Session(app.Env)

	return app
}
