package middleware

import (
	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(env *bootstrap.Env) gin.HandlerFunc {
	corsConfig := cors.Config{
		AllowOrigins:     []string{env.CorsOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
	}
	cors := cors.New(corsConfig)

	return cors
}
