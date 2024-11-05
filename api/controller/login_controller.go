package controller

import (
	"net/http"
	"strings"

	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	user, custErr := lc.LoginUsecase.GetUserByEmail(request.Email)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{Message: "invalid credentials"})
		return
	}

	accessToken, custErr := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	refreshToken, custErr := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	response := domain.LoginResponse{
		AccessToken: accessToken,
	}

	c.SetCookie("refresh_token", refreshToken, 3600*lc.Env.RefreshTokenExpiryHour, "/", strings.Split(c.Request.Host, ":")[0], true, true)

	c.JSON(http.StatusOK, response)
}
