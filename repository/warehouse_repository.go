package repository

import (
	"fmt"

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

func (wr *warehouseRepository) GetByID(warehouseID string) (domain.Warehouse, error) {
	var warehouse domain.Warehouse
	result := wr.database.Where("id = ?", warehouseID).Find(&warehouse)

	if result.RowsAffected == 0 {
		return warehouse, fmt.Errorf("no existing warehouse")
	}

	return warehouse, result.Error
}

func (wr *warehouseRepository) ModifyByID(warehouseID string, warehouse *domain.Warehouse) error {
	result := wr.database.Model(&domain.Warehouse{ID: warehouseID}).Updates(warehouse)

	return result.Error
}

func (wr *warehouseRepository) DeleteByID(warehouseID string) error {
	result := wr.database.Delete(&domain.Warehouse{}, "id = ?", warehouseID)

	return result.Error
}

func (wr *warehouseRepository) Fetch(name string) ([]domain.Warehouse, error) {
	var ws []domain.Warehouse
	query := wr.database.Model(&ws)

	if name != "" {
		query.Where("lower(name) LIKE lower(?)", "%"+name+"%")
	}

	result := query.Find(&ws)

	return ws, result.Error
}
