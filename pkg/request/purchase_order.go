package request

import (
	"errors"
	"github.com/miloalej-dev/W17-G1-Bootcamp/pkg/models"
	"net/http"
	"time"
)

type PurchaseOrderRequest struct {
	OrderNumber   *string               `json:"order_number"`
	OrderDate     *time.Time            `json:"order_date"`
	TracingCode   *string               `json:"tracing_code"`
	BuyersID      *int                  `json:"buyer_id"`
	WarehousesID  *int                  `json:"warehouse_id"`
	CarriersID    *int                  `json:"carrier_id"`
	OrderStatusID *int                  `json:"order_status_id"`
	OrderDetails  *[]models.OrderDetail `json:"order_details"`
}

func (p *PurchaseOrderRequest) Bind(r *http.Request) error {
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
