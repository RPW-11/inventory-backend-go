package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserRoute(group *gin.RouterGroup, db *gorm.DB) {
	ur := repository.NewUserRepository(db)

	uc := controller.UserController{
		UserUsecase: usecase.NewUserUsecase(ur),
	}

	group.GET("/profile", uc.GetUserProfile)
	group.GET("/user", uc.GetAllUsers)
}
