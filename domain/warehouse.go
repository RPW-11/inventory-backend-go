package domain

import "time"

const (
	TableWarehouse = "Warehouse"
)

type Warehouse struct {
	ID        string    `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Address   string    `gorm:"column:address" json:"address"`
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (Warehouse) TableName() string {
	return TableWarehouse
}

type WarehouseRepository interface {
	Create(warehouse *Warehouse) *CustomError
	GetByID(warehouseID string) (Warehouse, *CustomError)
	ModifyByID(warehouseID string, warehouse *Warehouse) *CustomError
	DeleteByID(warehouseID string) *CustomError
	Fetch(name string) ([]Warehouse, *CustomError)
}

type WarehouseUsecase interface {
	Create(warehouse *Warehouse) *CustomError
	ModifyByID(warehouseID string, warehouse *Warehouse) *CustomError
	DeleteByID(warehouseID string) *CustomError
	Fetch(name string) ([]Warehouse, *CustomError)
}
