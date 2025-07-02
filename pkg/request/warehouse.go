package request

import (
	"errors"
	"net/http"
)

// WharehouseRequest is a struct that represents a wharehouse in JSON format
type WarehouseRequest struct {
	Code               *string `json:"code"`
	Address            *string `json:"address"`
	Telephone          *string `json:"telephone"`
	MinimumCapacity    *int    `json:"minimum_capacity"`
	MinimumTemperature *int    `json:"minimum_temperature"`
}

func (p *WarehouseRequest) Bind(r *http.Request) error {
	if p.Code == nil {
		return errors.New("Code must not be null")
	}
	if p.Address == nil {
		return errors.New("Address must not be null")
	}
	if p.Telephone == nil {
		return errors.New("Telephone must not be null")
	}
	if p.MinimumCapacity == nil {
		return errors.New("Minimum Capacity must not be null")
	}
	if p.MinimumTemperature == nil {
		return errors.New("Minimum Temperature must not be null")
	}

	return nil
}
