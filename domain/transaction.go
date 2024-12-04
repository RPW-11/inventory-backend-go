package domain

import (
	"time"
)

const (
	TableTransaction = "Transaction"
	TransactionBuy   = "BUY"
	TransactionSell  = "SELL"
)

type Transaction struct {
	ID               int       `gorm:"primaryKey;column:id" json:"id"`
	ProductID        string    `gorm:"not null;column:product_id" json:"productId"`
	WarehouseID      string    `gorm:"not null;column:warehouse_id" json:"warehouseId"`
	Quantity         int       `gorm:"not null;check:quantity > 0;column:quantity" json:"quantity"`
	TotalPrice       int       `gorm:"not null;check:total_price >= 0;column:total_price" json:"totalPrice"`
	Description      string    `gorm:"type:text;column:description" json:"description"`
	EmployeeInCharge string    `gorm:"not null;column:employee_in_charge" json:"employeeInCharge"`
	TransactionType  string    `gorm:"not null;column:transaction_type" json:"transactionType"`
	PaymentStatus    string    `gorm:"default:'unpaid';column:payment_status" json:"paymentStatus"`
	CreatedAt        time.Time `gorm:"autoCreateTime;column:created_at" json:"createdAt"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime;column:updated_at" json:"updatedAt"`
}

func (Transaction) TableName() string {
	return TableTransaction
}

type TransactionRepository interface {
	Create(transaction *Transaction) *CustomError
}

type TransactionUseCase interface {
	Record(transaction *Transaction) *CustomError
}
