package repository

import (
	"net/http"

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

func (wr *warehouseRepository) Create(warehouse *domain.Warehouse) *domain.CustomError {
	result := wr.database.Create(warehouse)

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (wr *warehouseRepository) GetByID(warehouseID string) (domain.Warehouse, *domain.CustomError) {
	var warehouse domain.Warehouse
	result := wr.database.Where("id = ?", warehouseID).Find(&warehouse)

	if result.RowsAffected == 0 {
		return warehouse, domain.NewCustomError("warehouse doesnot exist", http.StatusBadRequest)
	}

	if result.Error != nil {
		return warehouse, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return warehouse, nil
}

func (wr *warehouseRepository) ModifyByID(warehouseID string, warehouse *domain.Warehouse) *domain.CustomError {
	result := wr.database.Model(&domain.Warehouse{ID: warehouseID}).Updates(warehouse)

	if result.RowsAffected == 0 {
		return domain.NewCustomError("invalid warehouse", http.StatusBadRequest)
	}

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (wr *warehouseRepository) DeleteByID(warehouseID string) *domain.CustomError {
	result := wr.database.Delete(&domain.Warehouse{}, "id = ?", warehouseID)

	if result.RowsAffected == 0 {
		return domain.NewCustomError("invalid warehouse", http.StatusBadRequest)
	}

	if result.Error != nil {
		return domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return nil
}

func (wr *warehouseRepository) Fetch(name string) ([]domain.Warehouse, *domain.CustomError) {
	var ws []domain.Warehouse
	query := wr.database.Model(&ws)

	if name != "" {
		query.Where("lower(name) LIKE lower(?)", "%"+name+"%")
	}

	result := query.Find(&ws)

	if result.Error != nil {
		return ws, domain.NewCustomError(result.Error.Error(), http.StatusInternalServerError)
	}

	return ws, nil
}
