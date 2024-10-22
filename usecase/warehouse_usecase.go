package usecase

import "github.com/RPW-11/inventory_management_be/domain"

type warehouseUsecase struct {
	warehouseRepository domain.WarehouseRepository
}

func NewWarehouseUsecase(wr domain.WarehouseRepository) domain.WarehouseUsecase {
	return &warehouseUsecase{
		warehouseRepository: wr,
	}
}

func (wu *warehouseUsecase) Create(warehouse *domain.Warehouse) error {
	return wu.warehouseRepository.Create(warehouse)
}

func (wu *warehouseUsecase) ModifyByID(warehouseID string, warehouse *domain.Warehouse) error {
	return wu.warehouseRepository.ModifyByID(warehouseID, warehouse)
}

func (wu *warehouseUsecase) DeleteByID(warehouseID string) error {
	return wu.warehouseRepository.DeleteByID(warehouseID)
}

func (wu *warehouseUsecase) Fetch(name string) ([]domain.Warehouse, error) {
	return wu.warehouseRepository.Fetch(name)
}
