package controller

import (
	"net/http"
	"time"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WarehouseController struct {
	WarehouseUsecase domain.WarehouseUsecase
}

func (wc *WarehouseController) CreateWarehouse(c *gin.Context) {
	var request domain.Warehouse
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	if request.Name == "" || request.Address == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "please fill all the required fields!"})
		return
	}

	warehouse := domain.Warehouse{
		ID:      uuid.NewString(),
		Name:    request.Name,
		Address: request.Address,
	}

	custErr := wc.WarehouseUsecase.Create(&warehouse)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Warehouse created successfully",
	})
}

func (wc *WarehouseController) ModifyWarehouseByID(c *gin.Context) {
	var request domain.Warehouse
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	wid := c.Param("id")

	if wid == "" || request.Name == "" || request.Address == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "please fill all the required fields! (id, name, address)"})
		return
	}

	warehouse := domain.Warehouse{
		Name:      request.Name,
		Address:   request.Address,
		CreatedAt: time.Now(),
	}

	custErr := wc.WarehouseUsecase.ModifyByID(wid, &warehouse)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Warehouse updated successfully",
	})
}

func (wc *WarehouseController) DeleteWarehouseByID(c *gin.Context) {
	wid := c.Param("id")
	if wid == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "no id given"})
		return
	}

	custErr := wc.WarehouseUsecase.DeleteByID(wid)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Warehouse deleted successfully",
	})
}

func (wc *WarehouseController) GetWarehouses(c *gin.Context) {
	warehouseName := c.Query("name")
	warehouses, custErr := wc.WarehouseUsecase.Fetch(warehouseName)
	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, warehouses)
}
