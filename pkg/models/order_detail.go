package models

type OrderDetail struct {
	ID               int      `gorm:"primaryKey;column:id"`
	Quantity         *int     `gorm:"column:quantity"`
	CleanLinesStatus *string  `gorm:"column:clean_lines_status;size:64"`
	Temperature      *float64 `gorm:"column:temperature;type:decimal(19,2)"`
	ProductRecordID  int      `gorm:"column:product_records_id;not null"`
	PurchaseOrderID  int      `gorm:"column:purchase_orders_id;not null"`
}
