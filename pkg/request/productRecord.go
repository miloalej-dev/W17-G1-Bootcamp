package request

import (
	"errors"
	"net/http"
)

type ProductRecordRequest struct {
	Id            *int     `json:"id"`
	LastUpdate    *string  `json:"last_update"`
	PurchasePrice *float64 `json:"purchase_price"`
	SalePrice     *float64 `json:"sale_price"`
	ProductsId    *int     `json:"products_id"`
}

func (b *ProductRecordRequest) Bind(r *http.Request) error {

	if b.LastUpdate == nil {
		return errors.New("Last update date must be not null")
	}
	if b.PurchasePrice == nil {
		return errors.New("Purchase price  must be not null")
	}
	if b.SalePrice == nil {
		return errors.New("Sale price must be not null")
	}
	if b.ProductsId == nil {
		return errors.New("Products Id must be not null")
	}

	return nil
}
