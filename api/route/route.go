package route

import (
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// register all the available routes
func Setup(env *bootstrap.Env, db *gorm.DB, gin *gin.Engine) {
	privateRouter := gin.Group("v1")

	// Private routes
	NewHomeRouter(privateRouter, db)

}
