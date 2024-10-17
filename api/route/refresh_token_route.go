package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRefreshTokenRoute(group *gin.RouterGroup, env *bootstrap.Env, db *gorm.DB) {
	rtu := usecase.NewRefreshTokenUsecase(repository.NewUserRepository(db))
	rtc := controller.RefreshTokenController{
		RefreshTokenUsecase: rtu,
		Env:                 env,
	}

	group.POST("/refresh", rtc.RefreshToken)
}
