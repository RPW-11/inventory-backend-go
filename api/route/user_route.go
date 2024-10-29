package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoute(group *gin.RouterGroup, db *gorm.DB, env *bootstrap.Env, storage *s3.S3) {
	ur := repository.NewUserRepository(db)
	sr := repository.NewStorageRepository(env, storage)

	uc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur, sr),
	}

	group.GET("/profile", uc.GetUserProfile)
	group.GET("/user", uc.GetAllUsers)
	group.PATCH("profile-picture", uc.UpdateProfilePicture)
}
