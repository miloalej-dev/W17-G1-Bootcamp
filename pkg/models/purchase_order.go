package models

import "time"

type PurchaseOrder struct {
	Id            int       `json:"id"`
	OrderNumber   string    `json:"order_number"`
	OrderDate     time.Time `json:"order_date"`
	TracingCode   string    `json:"tracing_code"`
	BuyerID       int       `json:"buyer_id"`
	WarehouseID   int       `json:"warehouse_id"`
	CarrierID     int       `json:"carrier_id"`
	OrderStatusID int       `json:"order_status_id"`
}
