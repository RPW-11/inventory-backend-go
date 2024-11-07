package controller

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/RPW-11/inventory_management_be/internal/fileutil"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

func (pc *ProductController) GetAllProducts(c *gin.Context) {
	productName := c.Query("name")

	products, custErr := pc.ProductUsecase.Fetch(productName)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) UploadProductImages(c *gin.Context) {
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB max memory
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "please provide product id"})
		return
	}

	_, custErr := pc.ProductUsecase.GetByID(productId)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	form := c.Request.MultipartForm
	imgFiles := form.File["product_imgs"]

	err = fileutil.CheckValidProductImages(imgFiles)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	custErr = pc.ProductUsecase.AddProductImages(imgFiles, productId)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "successfully add the product images"})
}

func (pc *ProductController) DeleteProductImage(c *gin.Context) {
	productImageId := c.Param("id")
	if productImageId == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "product's image id can't be empty"})
		return
	}

	custErr := pc.ProductUsecase.DeleteProductImage(productImageId)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "product's image has been deleted successfully"})
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "product's id can't be empty"})
		return
	}

	custErr := pc.ProductUsecase.DeleteProduct(productId)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "product has been deleted successfully"})
}
