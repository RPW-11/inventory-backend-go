package repository

import (
	"net/http"
	"time"

	"github.com/RPW-11/inventory_management_be/domain"
	"gorm.io/gorm"
)

type inventoryRepository struct {
	database *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) domain.InventoryRepository {
	return &inventoryRepository{
		database: db,
	}
}

func (ir *inventoryRepository) Create(inventory *domain.Inventory) *domain.CustomError {
	result := ir.database.Create(inventory)

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (ir *inventoryRepository) GetByID(id int) (domain.Inventory, *domain.CustomError) {
	var inventory domain.Inventory
	result := ir.database.Where("id = ?", id).Find(&inventory)

	if result.RowsAffected == 0 {
		return inventory, domain.NewCustomError("inventory doesnot exist", http.StatusBadRequest)
	}

	if result.Error != nil {
		return inventory, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return inventory, nil
}

func (ir *inventoryRepository) GetByProductWarehouseID(productID, warehouseID string) (domain.Inventory, *domain.CustomError) {
	var inventory domain.Inventory
	result := ir.database.Where(&domain.Inventory{ProductId: productID, WarehouseId: warehouseID}).Find(&inventory)

	if result.RowsAffected == 0 {
		return inventory, domain.NewCustomError("inventory does not exist", http.StatusBadRequest)
	}

	if result.Error != nil {
		return inventory, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return inventory, nil
}

func (ir *inventoryRepository) ModifyByID(inventoryID int, inventory *domain.Inventory) *domain.CustomError {
	inventory.UpdatedAt = time.Now()
	result := ir.database.Model(&domain.Inventory{ID: inventoryID}).Updates(inventory)

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (ir *inventoryRepository) GetByProductID(productID string) ([]domain.Inventory, *domain.CustomError) {
	var inventories []domain.Inventory
	result := ir.database.Where("product_id = ?", productID).Find(&inventories)

	if result.Error != nil {
		return inventories, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return inventories, nil
}
