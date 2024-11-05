package domain

import (
	"time"
)

const (
	TableInventory = "Inventory"
)

type Inventory struct {
	ID          int       `gorm:"column:id;primaryKey" json:"id"`
	ProductId   string    `gorm:"column:product_id" json:"productId"`
	WarehouseId string    `gorm:"column:warehouse_id" json:"warehouseId"`
	Quantity    int       `gorm:"column:quantity" json:"quantity"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

type WarehouseQuantity struct {
	WarehouseID     string `json:"warehouseId" binding:"required"`
	ProductQuantity int    `json:"productQuantity" binding:"required"`
}

type CreateInventoryRequest struct {
	ProductID          string              `json:"productId"`
	ProductName        string              `json:"productName" binding:"required"`
	ProductDescription string              `json:"productDescription" binding:"required"`
	ProductPrice       float64             `json:"productPrice" binding:"required"`
	Warehouses         []WarehouseQuantity `json:"warehouses" binding:"required"`
}

type CreateInventoryResponse struct {
	ProductID string `json:"productId"`
}

type UpdateQuantityRequest struct {
	Quantity int `json:"quantity" binding:"required"`
}

type ProductDetail struct {
	Product     Product           `json:"product"`
	Inventories []InventoryDetail `json:"inventories"`
	ImageUrls   []string          `json:"imageUrls"`
}

type InventoryDetail struct {
	ID               int    `json:"id"`
	WarehouseID      string `json:"warehouseId"`
	WarehouseName    string `json:"warehouseName"`
	WarehouseAddress string `json:"warehouseAddress"`
	ProductQuantity  int    `json:"productQuantity"`
}

func (Inventory) TableName() string {
	return TableInventory
}

type InventoryRepository interface {
	Create(inventory *Inventory) *CustomError
	GetByID(id int) (Inventory, *CustomError)
	GetByProductWarehouseID(productID, warehouseID string) (Inventory, *CustomError)
	ModifyByID(inventoryID int, inventory *Inventory) *CustomError
	GetByProductID(productID string) ([]Inventory, *CustomError)
}

type InventoryUsecase interface {
	CreateProductInventory(product *Product, warehouses []WarehouseQuantity) (string, *CustomError)
	GetByID(id int) (Inventory, *CustomError)
	GetProductDetails() ([]ProductDetail, *CustomError)
	ModifyByID(inventoryID int, inventory *Inventory) *CustomError
}
