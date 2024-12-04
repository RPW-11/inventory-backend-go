package controller

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionUsecase domain.TransactionUsecase
}

func (tc *TransactionController) RecordTransaction(c *gin.Context) {
	var request domain.TransactionRecordRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: err.Error()})
		return
	}

	// check quantityty
	if request.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, domain.Response{Message: "quantity must be not negative"})
		return
	}

	custErr := tc.TransactionUsecase.Record(&domain.Transaction{
		ProductID:        request.ProductID,
		WarehouseID:      request.WarehouseID,
		Quantity:         request.Quantity,
		Description:      request.Description,
		EmployeeInCharge: request.EmployeeInCharge,
		TransactionType:  request.TransactionType,
	})

	if custErr != nil {
		c.JSON(custErr.StatusCode, domain.Response{Message: custErr.Message})
		return
	}

	c.JSON(http.StatusOK, domain.Response{Message: "the transaction has been recorded successfully"})
}
