package domain

import "time"

const (
	TableInventory = "Inventory"
)

type Inventory struct {
	ID          int       `gorm:"column:id;primaryKey" json:"id"`
	ProductId   string    `gorm:"column:product_id" json:"productId"`
	WarehouseId string    `gorm:"column:warehouse_id" json:"warehouseId"`
	Quantity    int       `gorm:"column:quantity" json:"quantity"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

type CreateInventoryRequest struct {
	ProductID          string  `json:"productId"`
	ProductName        string  `json:"productName" binding:"required"`
	ProductDescription string  `json:"productDescription" binding:"required"`
	ProductPrice       float64 `json:"productPrice" binding:"required"`
	ProductQuantity    int     `json:"productQuantity" binding:"required"`
	WarehouseID        string  `json:"warehouseId" binding:"required"`
}

func (Inventory) TableName() string {
	return TableInventory
}

type InventoryRepository interface {
	Create(inventory *Inventory) error
	GetByProductWarehouseID(productID, warehouseID string) (Inventory, error)
	ModifyByID(inventoryID int, inventory *Inventory) error
}

type InventoryUsecase interface {
	CreateProductInventory(product *Product, warehouseID string, quantity int) error
}
