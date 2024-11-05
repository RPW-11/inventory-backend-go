package middleware

import (
	"net/http"
	"strings"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		position := c.GetString("x-user-position")

		if strings.ToLower(position) != "admin" {
			c.JSON(http.StatusUnauthorized, domain.Response{Message: "you are not an admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
