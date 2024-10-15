package controller

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var request domain.Product

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	if request.Name == "" || request.Price <= 0 || request.Description == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "Please fill all the required fields!"})
		return
	}

	product := domain.Product{
		ID:          uuid.NewString(),
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}

	err = pc.ProductUsecase.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Product created successfully",
	})
}

func (pc *ProductController) GetProductByID(c *gin.Context) {

}

func (pc *ProductController) GetProducts(c *gin.Context) {

}
