package route

import (
	"github.com/RPW-11/inventory_management_be/api/middleware"
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// register all the available routes
func Setup(env *bootstrap.Env, db *gorm.DB, storage *s3.S3, gin *gin.Engine) {
	// CORS setup
	cors := middleware.CorsMiddleware(env)

	gin.Use(cors)

	router := gin.Group("v1")

	// Public routes
	publicRouter := router.Group("")

	NewSignupRouter(publicRouter, env, db)
	NewLoginRoute(publicRouter, env, db)
	NewRefreshTokenRoute(publicRouter, env, db)

	// Private routes
	jwtMiddleware := middleware.JwtAuthMiddleware(env.AccessTokenSecret)
	privateRouter := router.Group("")
	privateRouter.Use(jwtMiddleware)
	NewInventoryRoute(privateRouter, env, db)
	NewWarehouseRoute(privateRouter, db)
	NewProductRoute(privateRouter, db)
	NewUserRoute(privateRouter, db, env, storage)
}
