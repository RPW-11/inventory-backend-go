package middleware

import (
	"net/http"
	"strings"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/RPW-11/inventory_management_be/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]

			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, userPosition, err := tokenutil.ExtractPositionIDFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.Response{Message: err.Error()})
					c.Abort()
					return
				}

				c.Set("x-user-id", userID)
				c.Set("x-user-position", userPosition)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.Response{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.Response{Message: "Not authorized"})
		c.Abort()
	}
}
