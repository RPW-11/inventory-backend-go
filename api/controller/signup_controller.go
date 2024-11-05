package controller

import (
	"net/http"
	"strings"

	"github.com/RPW-11/inventory_management_be/bootstrap"
	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	_, custErr := sc.SignupUsecase.GetUserByEmail(request.Email)
	if custErr == nil {
		c.JSON(http.StatusConflict, domain.Response{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:          uuid.NewString(),
		FullName:    request.FullName,
		Email:       request.Email,
		Position:    "Staff",
		Password:    request.Password,
		PhoneNumber: request.PhoneNumber,
	}

	custErr = sc.SignupUsecase.Create(&user)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	accessToken, custErr := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	refreshToken, custErr := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	response := domain.SignupResponse{
		AccessToken: accessToken,
	}

	c.SetCookie("refresh_token", refreshToken, 3600*sc.Env.RefreshTokenExpiryHour, "/", strings.Split(c.Request.Host, ":")[0], true, true)

	c.JSON(http.StatusOK, response)
}
