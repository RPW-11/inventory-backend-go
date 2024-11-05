package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/api/middleware"
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewProductRoute(group *gin.RouterGroup, env *bootstrap.Env, db *gorm.DB, storage *s3.S3) {
	pr := repository.NewProductRepository(db)
	sr := repository.NewStorageRepository(env, storage)

	pc := controller.ProductController{
		ProductUsecase: usecase.NewProductUsecase(pr, sr),
	}

	roleMiddleware := middleware.RoleMiddleware()

	group.GET("/product", pc.GetAllProducts)
	group.POST("/product-images/:id", roleMiddleware, pc.UploadProductImages)
	group.DELETE("/product-image/:id", roleMiddleware, pc.DeleteProductImage)
}
