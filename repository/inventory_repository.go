package repository

import (
	"fmt"

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

func (ir *inventoryRepository) Create(inventory *domain.Inventory) error {
	result := ir.database.Create(inventory)

	return result.Error
}

func (ir *inventoryRepository) GetByProductWarehouseID(productID, warehouseID string) (domain.Inventory, error) {
	var inventory domain.Inventory
	result := ir.database.Where(&domain.Inventory{ProductId: productID, WarehouseId: warehouseID}).Find(&inventory)

	if result.RowsAffected == 0 {
		return inventory, fmt.Errorf("no existing inventory")
	}

	return inventory, result.Error
}

func (ir *inventoryRepository) ModifyByID(inventoryID int, inventory *domain.Inventory) error {
	result := ir.database.Model(&domain.Inventory{ID: inventoryID}).Updates(inventory)

	return result.Error
}

func (ir *inventoryRepository) GetByProductID(productID string) ([]domain.Inventory, error) {
	var inventories []domain.Inventory
	result := ir.database.Where("product_id = ?", productID).Find(&inventories)

	return inventories, result.Error
}
