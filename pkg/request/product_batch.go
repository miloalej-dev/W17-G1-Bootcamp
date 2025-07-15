package request

import (
	"errors"
	"net/http"
)

type ProductBatchRequest struct {
	Id                 *int     `json:"id"`
	BatchNumber        *int     `json:"batch_number"`
	CurrentQuantity    *int     `json:"current_quantity"`
	CurrentTemperature *float64 `json:"current_temperature"`
	DueDate            *string  `json:"due_date"`
	InitialQuantity    *int     `json:"initial_quantity"`
	ManufacturingDate  *string  `json:"manufacturing_date"`
	ManufacturingHour  *int     `json:"manufacturing_hour"`
	MinumumTemperature *float64 `json:"minumum_temperature"`
	SectionsId         *int     `json:"sections_id"`
	ProductsId         *int     `json:"products_id"`
}

func (p *ProductBatchRequest) Bind(r *http.Request) error {
	if p.BatchNumber == nil {
		return errors.New("BatchNumber must not be null")
	}
	if p.CurrentQuantity == nil {
		return errors.New("CurrentQuantity must not be null")
	}
	if p.CurrentTemperature == nil {
		return errors.New("CurrentTemperature must not be null")
	}
	if p.DueDate == nil {
		return errors.New("DueDate must not be null")
	}
	if p.InitialQuantity == nil {
		return errors.New("InitialQuantity must not be null")
	}
	if p.ManufacturingDate == nil {
		return errors.New("ManufacturingDate must not be null")
	}
	if p.ManufacturingHour == nil {
		return errors.New("ManufacturingHour must not be null")
	}
	if p.MinumumTemperature == nil {
		return errors.New("MinumumTemperature must not be null")
	}
	if p.SectionsId == nil {
		return errors.New("SectionsId must not be null")
	}
	if p.ProductsId == nil {
		return errors.New("ProductsId must not be null")
	}

	return nil
}
