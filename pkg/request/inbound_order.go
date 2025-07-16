package request

import (
	"errors"
	"net/http"
)

type InboundOrder struct {
	Id             *int    `json:"id"`
	OrderNumber    *string `json:"order_number"`
	EmployeeId     *int    `json:"employee_id"`
	ProductBatchId *int    `json:"product_batch_id"`
	WarehouseId    *int    `json:"warehouse_id"`
}

func (i *InboundOrder) Bind(r *http.Request) error {
	if i.OrderNumber == nil {
		return errors.New("order_number must not be null")
	}
	if i.EmployeeId == nil {
		return errors.New("employee_id must not be null")
	}
	if i.ProductBatchId == nil {
		return errors.New("product_batch_id must not be null")
	}
	if i.WarehouseId == nil {
		return errors.New("warehouse_id must not be null")
	}
	return nil
}
