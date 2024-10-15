package repository

import (
	"github.com/RPW-11/inventory_management_be/domain"
	"gorm.io/gorm"
)

type warehouseRepository struct {
	database *gorm.DB
}

func NewWarehouseRepository(db *gorm.DB) domain.WarehouseRepository {
	return &warehouseRepository{
		database: db,
	}
}

func (wr *warehouseRepository) Create(warehouse *domain.Warehouse) error {
	result := wr.database.Create(warehouse)

	return result.Error
}

func (wr *warehouseRepository) ModifyByID(warehouseID string, warehouse *domain.Warehouse) error {
	result := wr.database.Model(&domain.Warehouse{ID: warehouseID}).Updates(warehouse)

	return result.Error
}

func (wr *warehouseRepository) DeleteByID(warehouseID string) error {
	result := wr.database.Delete(&domain.Warehouse{}, "id = ?", warehouseID)

	return result.Error
}

func (wr *warehouseRepository) Fetch() ([]domain.Warehouse, error) {
	var ws []domain.Warehouse
	result := wr.database.Find(&ws)

	return ws, result.Error
}
