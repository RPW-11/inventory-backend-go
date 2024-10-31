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

	products, err := pc.ProductUsecase.Fetch(productName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
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

	_, err = pc.ProductUsecase.GetByID(productId)
	if err != nil {
		if err.Error() == "no existing product" {
			c.JSON(http.StatusBadRequest, domain.Response{Message: "invalid product id"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
			return
		}
	}

	form := c.Request.MultipartForm
	imgFiles := form.File["product_imgs"]

	err = fileutil.CheckValidProductImages(imgFiles)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	err = pc.ProductUsecase.AddProductImages(imgFiles, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "successfully add the product images"})
}

func (pc *ProductController) DeleteProductImage(c *gin.Context) {
	productImageId := c.Param("id")
	if productImageId == "" {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: "product's image id can't be empty"})
		return
	}

	err := pc.ProductUsecase.DeleteProductImage(productImageId)
	if err != nil {
		if err.Error() == "no existing product's image" {
			c.JSON(http.StatusBadRequest, domain.Response{Message: "product's image doesn't exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "product's image has been deleted successfully"})
}
