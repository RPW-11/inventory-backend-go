package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/api/middleware"
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

	roleMiddleware := middleware.RoleMiddleware()

	group.GET("/product-inventory", ic.GetProductDetails)
	group.GET("/product-inventory/:id", ic.GetSingleProductDetail)

	group.POST("/inventory", roleMiddleware, ic.CreateProductInventory)
	group.PATCH("/inventory-update-quantity/:id", roleMiddleware, ic.UpdateQuantity)
}
