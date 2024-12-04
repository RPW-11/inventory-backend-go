package repository

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
	"gorm.io/gorm"
)

type transactionRepository struct {
	database *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{
		database: db,
	}
}

func (tr *transactionRepository) Create(transaction *domain.Transaction) *domain.CustomError {
	result := tr.database.Create(transaction)

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}
