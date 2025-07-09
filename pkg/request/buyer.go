package request

import (
	"errors"
	"net/http"
)

type BuyerRequest struct {
	CardNumberId *string `json:"card_number_id"`
	FirstName    *string `json:"first_name"`
	LastName     *string `json:"last_name"` // Telephone is the telephone of the seller company
}

func (b *BuyerRequest) Bind(r *http.Request) error {
	if b.CardNumberId == nil {
		return errors.New("Card number Id must be not null")
	}
	if b.FirstName == nil {
		return errors.New("Fisrt name must not be null")
	}
	if b.LastName == nil {
		return errors.New("Last name must not be null")
	}
	return nil
}
