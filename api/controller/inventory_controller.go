package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	InventoryUsecase domain.InventoryUsecase
}

func (ic *InventoryController) CreateProductInventory(c *gin.Context) {
	var request domain.CreateInventoryRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	if request.ProductPrice <= 0 {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "product price must be more than 0"})
		return
	}

	for _, warehouse := range request.Warehouses {
		if warehouse.ProductQuantity <= 0 {
			c.JSON(http.StatusBadRequest, domain.Response{Message: "product quantity must be more than 0"})
			return
		}
	}

	product := domain.Product{
		ID:          request.ProductID,
		Name:        request.ProductName,
		Description: request.ProductDescription,
		Price:       request.ProductPrice,
	}

	productId, custErr := ic.InventoryUsecase.CreateProductInventory(&product, request.Warehouses)

	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.CreateInventoryResponse{
		ProductID: productId,
	})
}

func (ic *InventoryController) GetProductDetails(c *gin.Context) {
	productName := c.Query("name")
	pageSize, offset := 10, 0

	if c.Query("pageSize") != "" {
		val, err := strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.Response{Message: "invalid page size"})
		}

		pageSize = val
	}

	if c.Query("offset") != "" {
		val, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.JSON(http.StatusBadRequest, domain.Response{Message: "invalid offset"})
		}

		offset = val
	}

	productDetails, custErr := ic.InventoryUsecase.GetProductDetails(productName, pageSize, offset)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, productDetails)
}

func (ic *InventoryController) GetSingleProductDetail(c *gin.Context) {
	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "please provide the product's id"})
		return
	}

	productDetail, custErr := ic.InventoryUsecase.GetProductDetailByID(productId)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, productDetail)
}

func (ic *InventoryController) UpdateQuantity(c *gin.Context) {
	var request domain.UpdateQuantityRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	// check if the id is valid
	invID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "invalid inventory id"})
		return
	}

	existing, custErr := ic.InventoryUsecase.GetByID(invID)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	// check if quantity valid
	if request.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "invalid quantity"})
		return
	}

	// update the quantity
	existing.Quantity = request.Quantity
	existing.ID = 0
	existing.UpdatedAt = time.Now()

	custErr = ic.InventoryUsecase.ModifyByID(invID, &existing)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Quantity updated successfully"})
}
