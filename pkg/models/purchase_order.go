package models

import "time"

type PurchaseOrder struct {
	Id            int       `json:"id"`
	OrderNumber   string    `json:"order_number"`
	OrderDate     time.Time `json:"order_date"`
	TracingCode   string    `json:"tracing_code"`
	BuyersID      int       `json:"buyers_id"`
	WarehousesID  int       `json:"warehouses_id"`
	CarriersID    int       `json:"carriers_id"`
	OrderStatusID int       `json:"order_status_id"`
}
