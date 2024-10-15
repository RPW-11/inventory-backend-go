package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewWarehouseRoute(group *gin.RouterGroup, db *gorm.DB) {
	wr := repository.NewWarehouseRepository(db)
	wc := controller.WarehouseController{
		WarehouseUsecase: usecase.NewWarehouseUsecase(wr),
	}
	group.GET("/warehouse", wc.GetWarehouses)
	group.POST("/warehouse", wc.CreateWarehouse)
	group.PUT("/warehouse/:id", wc.ModifyWarehouseByID)
	group.DELETE("/warehouse/:id", wc.DeleteWarehouseByID)
}
