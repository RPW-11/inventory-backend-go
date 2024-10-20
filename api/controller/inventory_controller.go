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

	if request.ProductPrice <= 0 || request.ProductQuantity < 0 {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "Invalid number of price or quantity"})
		return
	}

	product := domain.Product{
		ID:          request.ProductID,
		Name:        request.ProductName,
		Description: request.ProductDescription,
		Price:       request.ProductPrice,
	}

	err = ic.InventoryUsecase.CreateProductInventory(&product, request.WarehouseID, request.ProductQuantity)
	if err == nil {
		c.JSON(http.StatusOK, domain.Response{
			Message: "Inventory created successfully",
		})
		return
	}

	switch err.Error() {
	case "no existing warehouse":
		c.JSON(http.StatusBadRequest, domain.Response{Message: "Invalid warehouse"})
		return
	default:
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}
}

func (ic *InventoryController) GetProductDetails(c *gin.Context) {
	productDetails, err := ic.InventoryUsecase.GetProductDetails()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, productDetails)
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

	existing, err := ic.InventoryUsecase.GetByID(invID)
	if err != nil {
		if err.Error() == "no existing inventory" {
			c.JSON(http.StatusBadRequest, domain.Response{Message: "invalid inventory id"})
			return
		}
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
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

	err = ic.InventoryUsecase.ModifyByID(invID, &existing)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "Quantity updated successfully"})
}
