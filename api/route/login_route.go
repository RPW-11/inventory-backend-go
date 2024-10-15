package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLoginRoute(group *gin.RouterGroup, env *bootstrap.Env, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	lc := controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur),
		Env:          env,
	}

	group.POST("/login", lc.Login)
}
