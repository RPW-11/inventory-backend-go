package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductRoute(group *gin.RouterGroup, db *gorm.DB) {
	pr := repository.NewProductRepository(db)

	pc := controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr),
	}

	group.GET("/product", pc.GetAllProducts)
}
