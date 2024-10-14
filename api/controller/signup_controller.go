package controller

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.Response{Message: "User already exists with the given email"})
		return
	}

	user := domain.User{
		ID:       uuid.NewString(),
		FullName: request.FullName,
		Email:    request.Email,
		Position: "Staff",
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Success"})

}
