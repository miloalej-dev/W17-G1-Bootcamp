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
		return errors.New("Quantity must not be null")
	}
	if o.CleanLinesStatus == nil {
		return errors.New("CleanLinesStatus must not be null")
	}
	if o.Temperature == nil {
		return errors.New("Temperature must not be null")
	}
	if o.ProductRecordID == nil {
		return errors.New("ProductRecordID must not be null")
	}
	if o.PurchaseOrderID == nil {
		return errors.New("PurchaseOrderID must not be null")
	}
	return nil
}
