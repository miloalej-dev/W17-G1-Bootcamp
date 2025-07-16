package request

import (
	"errors"
	"net/http"
)

// WharehouseRequest is a struct that represents a wharehouse in JSON format
type WarehouseRequest struct {
	WarehouseCode		*string `json:"warehouse_code"`
	Address				*string `json:"address"`
	Telephone			*string `json:"telephone"`
	MinimumCapacity		*int    `json:"minimum_capacity"`
	MinimumTemperature	*int    `json:"minimum_temperature"`
	LocalityId			*int	`json:"locality_id"`
}

func (p *WarehouseRequest) Bind(r *http.Request) error {
	if p.WarehouseCode == nil {
		return errors.New("warehouse code must not be null")
	}
	if p.Address == nil {
		return errors.New("address must not be null")
	}
	if p.Telephone == nil {
		return errors.New("telephone must not be null")
	}
	if p.MinimumCapacity == nil {
		return errors.New("minimum Capacity must not be null")
	}
	if p.MinimumTemperature == nil {
		return errors.New("minimum Temperature must not be null")
	}
	if p.LocalityId == nil {
		return errors.New("locality id  must not be null")
	}

	return nil
}
