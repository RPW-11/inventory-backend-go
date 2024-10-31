package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewInventoryRoute(group *gin.RouterGroup, env *bootstrap.Env, db *gorm.DB) {
	pr := repository.NewProductRepository(db)
	ir := repository.NewInventoryRepository(db)
	wr := repository.NewWarehouseRepository(db)

	ic := controller.InventoryController{
		InventoryUsecase: usecase.NewInventoryUsecase(ir, pr, wr),
	}

	group.GET("/product-inventory", ic.GetProductDetails)
	group.POST("/inventory", ic.CreateProductInventory)
	group.PATCH("/inventory-update-quantity/:id", ic.UpdateQuantity)
}
