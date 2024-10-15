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
		c.JSON(http.StatusBadRequest, domain.Response{Message: "Please fill all the required fields!"})
		return
	}

	warehouse := domain.Warehouse{
		ID:      uuid.NewString(),
		Name:    request.Name,
		Address: request.Address,
	}

	err = wc.WarehouseUsecase.Create(&warehouse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
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
		c.JSON(http.StatusBadRequest, domain.Response{Message: "Please fill all the required fields! (id, name, address)"})
		return
	}

	warehouse := domain.Warehouse{
		Name:      request.Name,
		Address:   request.Address,
		CreatedAt: time.Now(),
	}

	err = wc.WarehouseUsecase.ModifyByID(wid, &warehouse)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Warehouse updated successfully",
	})
}

func (wc *WarehouseController) DeleteWarehouseByID(c *gin.Context) {
	wid := c.Param("id")
	if wid == "" {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "No id given"})
		return
	}

	err := wc.WarehouseUsecase.DeleteByID(wid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Message: "Warehouse deleted successfully",
	})
}

func (wc *WarehouseController) GetWarehouses(c *gin.Context) {
	warehouses, err := wc.WarehouseUsecase.Fetch()
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, warehouses)
}
