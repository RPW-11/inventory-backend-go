package controller

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	currentRefreshToken, err := c.Cookie("refresh_token")

	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{Message: err.Error()})
		return
	}

	id, _, err := rtc.RefreshTokenUsecase.ExtractPositionIDFromToken(currentRefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{Message: "Invalid refresh token"})
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{Message: "User not found"})
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	response := domain.RefreshTokenResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, response)
}
