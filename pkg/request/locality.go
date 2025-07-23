package request

import (
	"errors"
	"net/http"
)

type LocalityRequest struct {
	Id       int     `json:"id"` // El ID debe ser proporcionado por el cliente
	Locality *string `json:"locality_name"`
	Province *string `json:"province_name,omitempty"`
	Country  *string `json:"country_name,omitempty"`
}

func (l *LocalityRequest) Bind(r *http.Request) error {
	if l.Id <= 0 {
		return errors.New("locality_id must be greater than 0")
	}
	if l.Locality == nil {
		return errors.New("locality_name")
	}
	if l.Province == nil {
		return errors.New("province_name")
	}
	if l.Country == nil {
		return errors.New("country_name")
	}
	return nil
}
