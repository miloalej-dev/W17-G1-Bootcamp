package request

import (
	"errors"
	"net/http"
)

type SectionRequest struct {
	SectionNumber      *string  `json:"section_number"`
	CurrentTemperature *float64 `json:"current_temperature"`
	MinimumTemperature *float64 `json:"minimum_temperature"`
	CurrentCapacity    *int     `json:"current_capacity"`
	MinimumCapacity    *int     `json:"minimum_capacity"`
	MaximumCapacity    *int     `json:"maximum_capacity"`
	WarehousesId       *int     `json:"warehouses_id"`
	ProductTypeId      *int     `json:"product_type_id"`
}

func (p *SectionRequest) Bind(r *http.Request) error {
	if p.SectionNumber == nil {
		return errors.New("section_number is required")
	}
	if p.CurrentTemperature == nil {
		return errors.New("current_temperature is required")
	}
	if p.MinimumTemperature == nil {
		return errors.New("minimum_temperature is required")
	}
	if p.CurrentCapacity == nil {
		return errors.New("current_capacity is required")
	}
	if p.MinimumCapacity == nil {
		return errors.New("minimum_capacity is required")
	}
	if p.MaximumCapacity == nil {
		return errors.New("maximum_capacity is required")
	}

	return nil
}
