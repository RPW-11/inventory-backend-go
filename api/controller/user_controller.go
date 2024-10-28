package controller

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func (uc *UserController) GetUserProfile(c *gin.Context) {
	userId := c.GetString("x-user-id")

	profile, err := uc.UserUsecase.GetProfile(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.UserUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
