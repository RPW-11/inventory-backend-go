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

func (wu *warehouseUsecase) Create(warehouse *domain.Warehouse) *domain.CustomError {
	return wu.warehouseRepository.Create(warehouse)
}

func (wu *warehouseUsecase) ModifyByID(warehouseID string, warehouse *domain.Warehouse) *domain.CustomError {
	return wu.warehouseRepository.ModifyByID(warehouseID, warehouse)
}

func (wu *warehouseUsecase) DeleteByID(warehouseID string) *domain.CustomError {
	return wu.warehouseRepository.DeleteByID(warehouseID)
}

func (wu *warehouseUsecase) Fetch(name string) ([]domain.Warehouse, *domain.CustomError) {
	return wu.warehouseRepository.Fetch(name)
}
