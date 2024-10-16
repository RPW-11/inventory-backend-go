package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewInventoryRoute(group *gin.RouterGroup, db *gorm.DB) {
	pr := repository.NewProductRepository(db)
	ir := repository.NewInventoryRepository(db)
	wr := repository.NewWarehouseRepository(db)

	ic := controller.InventoryController{
		InventoryUsecase: usecase.NewInventoryUsecase(ir, pr, wr),
	}

	group.POST("/inventory", ic.CreateProductInventory)
}
