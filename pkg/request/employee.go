package request

import (
	"errors"
	"net/http"
)

type EmployeeRequest struct {
	Id           *int    `json:"id"`
	CardNumberId *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"`
	WarehouseId  *int    `json:"warehouse_id"`
}

func (p *EmployeeRequest) Bind(r *http.Request) error {
	if p.CardNumberId == nil {
		return errors.New("CardNumberId must not be null")
	}
	if p.FirstName == nil {
		return errors.New("FirstName must not be null")
	}
	if p.LastName == nil {
		return errors.New("LastName must not be null")
	}
	if p.WarehouseId == nil {
		return errors.New("WarehouseId must not be null")
	}
	return nil
}
