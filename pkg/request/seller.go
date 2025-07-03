package request

import (
	"errors"
	"net/http"
)

type SellerRequest struct {
	Name      *string `json:"name"`      // Name is the name of the seller company
	Address   *string `json:"address"`   // Address is the address of the seller company
	Telephone *string `json:"telephone"` // Telephone is the telephone of the seller company
}

func (p *SellerRequest) Bind(r *http.Request) error {
	if p.Name == nil {
		return errors.New("name must not be null")
	}
	if p.Address == nil {
		return errors.New("address must not be null")
	}
	if p.Telephone == nil {
		return errors.New("telephone must not be null")
	}
	return nil
}
