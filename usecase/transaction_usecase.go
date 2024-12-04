package usecase

import (
	"net/http"

	"github.com/RPW-11/inventory_management_be/domain"
)

type transactionUsecase struct {
	transactionRepository domain.TransactionRepository
	inventoryRepository   domain.InventoryRepository
	userRepository        domain.UserRepository
	productRepository     domain.ProductRepository
}

func NewTransactionUsecase(tr domain.TransactionRepository, ir domain.InventoryRepository, ur domain.UserRepository, pr domain.ProductRepository) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepository: tr,
		inventoryRepository:   ir,
		userRepository:        ur,
		productRepository:     pr,
	}
}

func (tu *transactionUsecase) Record(transaction *domain.Transaction) *domain.CustomError {
	// get the inventory from which the product and the warehouse in the transaction
	inventory, err := tu.inventoryRepository.GetByProductWarehouseID(transaction.ProductID, transaction.WarehouseID)
	if err != nil {
		return err
	}

	// check the employee in charge
	_, err = tu.userRepository.GetByID(transaction.EmployeeInCharge)
	if err != nil {
		return err
	}

	// get the product's price
	product, err := tu.productRepository.GetByID(transaction.ProductID)
	if err != nil {
		return err
	}

	// set the total price in the transaction
	transaction.TotalPrice = product.Price * float64(transaction.Quantity)

	// check transaction type
	if transaction.TransactionType == domain.TransactionBuy {
		inventory.Quantity += transaction.Quantity
	} else if transaction.TransactionType == domain.TransactionSell {
		// decrement the quantity value of the product
		if transaction.Quantity > inventory.Quantity {
			return domain.NewCustomError("not enough quantity in the inventory", http.StatusBadRequest)
		}
		inventory.Quantity -= transaction.Quantity
	} else {
		return domain.NewCustomError("invalid transaction type", http.StatusBadRequest)
	}

	// update the inventory
	err = tu.inventoryRepository.ModifyByID(inventory.ID, &inventory)
	if err != nil {
		return err
	}

	// insert the transaction record
	err = tu.transactionRepository.Create(transaction)
	if err != nil {
		return err
	}

	return nil
}
