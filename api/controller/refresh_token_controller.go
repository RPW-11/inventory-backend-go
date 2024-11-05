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

	id, _, custErr := rtc.RefreshTokenUsecase.ExtractPositionIDFromToken(currentRefreshToken, rtc.Env.RefreshTokenSecret)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	user, custErr := rtc.RefreshTokenUsecase.GetUserByID(id)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	accessToken, custErr := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	response := domain.RefreshTokenResponse{
		AccessToken: accessToken,
	}

	c.JSON(http.StatusOK, response)
}
