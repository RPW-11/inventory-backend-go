package controller

import (
	"net/http"
	"strings"

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

func (uc *UserController) UpdateProfilePicture(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("profile_img")
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	defer file.Close()

	contentType := fileHeader.Header.Get("Content-Type")
	if contentType != "image/png" && contentType != "image/jpeg" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "invalid file type"})
		return
	}

	if fileHeader.Size > 5000000 {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "image must be less than 5 MB"})
		return
	}

	userId := c.GetString("x-user-id")
	fileHeader.Filename = userId + "." + strings.Split(contentType, "/")[1]

	err = uc.UserUsecase.UpdateProfilePicture(userId, file, fileHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "profile updated successfully"})
}
