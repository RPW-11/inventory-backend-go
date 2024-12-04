package route

import (
	"github.com/RPW-11/inventory_management_be/api/controller"
	"github.com/RPW-11/inventory_management_be/repository"
	"github.com/RPW-11/inventory_management_be/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTransactionRouter(group *gin.RouterGroup, db *gorm.DB) {
	ur := repository.NewUserRepository(db)
	pr := repository.NewProductRepository(db)
	ir := repository.NewInventoryRepository(db)
	tr := repository.NewTransactionRepository(db)

	tc := controller.TransactionController{
		TransactionUsecase: usecase.NewTransactionUsecase(tr, ir, ur, pr),
	}

	group.POST("/record-transaction", tc.RecordTransaction)
}
