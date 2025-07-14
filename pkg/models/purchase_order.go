package models

import "time"

type PurchaseOrder struct {
	ID            int       `gorm:"primaryKey;column:id"`
	OrderNumber   string    `gorm:"column:order_number;size:64"`
	OrderDate     time.Time `gorm:"column:order_date"`
	TracingCode   string    `gorm:"column:tracing_code;size:64"`
	BuyersID      int       `gorm:"column:buyers_id;not null"`
	WarehousesID  int       `gorm:"column:warehouses_id;not null"`
	CarriersID    int       `gorm:"column:carriers_id;not null"`
	OrderStatusID int       `gorm:"column:order_status_id;not null"`
}
