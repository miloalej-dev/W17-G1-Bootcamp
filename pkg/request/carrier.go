package request

import (
	"errors"
	"net/http"
)

// WharehouseRequest is a struct that represents a wharehouse in JSON format
type CarrierRequest struct {
	ID				*int	`json:"id"`
	CId				*string	`json:"cid" gorm:"column:cid"`
	CompanyName		*string	`json:"company_name" gorm:"column:name"`
	Address			*string	`json:"address"`
	Telephone		*string	`json:"telephone"`
	LocalityId		*int	`json:"locality_id"`
}

func (p *CarrierRequest) Bind(r *http.Request) error {
	if p.CId == nil {
		return errors.New("cid code must not be null")
	}
	if p.CompanyName == nil {
		return errors.New("company name must not be null")
	}
	if p.Address == nil {
		return errors.New("address must not be null")
	}
	if p.Telephone == nil {
		return errors.New("telephone must not be null")
	}
	if p.LocalityId == nil {
		return errors.New("locality id must not be null")
	}

	return nil
}
