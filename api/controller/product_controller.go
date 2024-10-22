package controller

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	productName := c.Query("name")

	products, err := pc.ProductUsecase.Fetch(productName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
