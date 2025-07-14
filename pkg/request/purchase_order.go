package request

import (
	"errors"
	"net/http"
	"time"
)

type PurchaseOrderRequest struct {
	ID            *int       `json:"id"`
	OrderNumber   *string    `json:"order_number"`
	OrderDate     *time.Time `json:"order_date"`
	TracingCode   *string    `json:"tracing_code"`
	BuyersID      *int       `json:"buyers_id"`
	WarehousesID  *int       `json:"warehouses_id"`
	CarriersID    *int       `json:"carriers_id"`
	OrderStatusID *int       `json:"order_status_id"`
}

func (p *PurchaseOrderRequest) Bind(r *http.Request) error {
	if p.ID == nil {
		return errors.New("ID must not be null")
	}
	if p.OrderNumber == nil {
		return errors.New("OrderNumber must not be null")
	}
	if p.OrderDate == nil {
		return errors.New("OrderDate must not be null")
	}
	if p.TracingCode == nil {
		return errors.New("TracingCode must not be null")
	}
	if p.BuyersID == nil {
		return errors.New("BuyersID must not be null")
	}
	if p.WarehousesID == nil {
		return errors.New("WarehousesID must not be null")
	}
	if p.CarriersID == nil {
		return errors.New("CarriersID must not be null")
	}
	if p.OrderStatusID == nil {
		return errors.New("OrderStatusID must not be null")
	}
	return nil
}
