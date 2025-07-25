package request

import (
	"errors"
	"net/http"
)

type OrderDetailRequest struct {
	Quantity         *int     `json:"quantity"`
	CleanLinesStatus *string  `json:"clean_lines_status"`
	Temperature      *float64 `json:"temperature"`
	ProductRecordID  *int     `json:"product_records_id"`
	PurchaseOrderID  *int     `json:"purchase_orders_id"`
}

func (o *OrderDetailRequest) Bind(r *http.Request) error {

	if o.Quantity == nil {
		return errors.New("quantity must not be null")
	}
	if o.CleanLinesStatus == nil {
		return errors.New("clean line status must not be null")
	}
	if o.Temperature == nil {
		return errors.New("temperature must not be null")
	}
	if o.ProductRecordID == nil {
		return errors.New("product record Id must not be null")
	}
	if o.PurchaseOrderID == nil {
		return errors.New("purchase order Id must not be null")
	}
	return nil
}
